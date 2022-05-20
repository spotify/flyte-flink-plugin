package flink

import (
	"testing"

	flinkOp "github.com/spotify/flink-on-k8s-operator/apis/flinkcluster/v1beta1"
	flinkIdl "github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink"
	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var exampleArtifacts = []string{"gs://bucket/a.jar", "gs://bucket/b.jar", "gs://bucket/c.jar"}

func TestNewFlinkCluster(t *testing.T) {

	tests := []struct {
		name    string
		taskCtx *FlinkTaskContext
		check   func(t *testing.T, cluster *flinkOp.FlinkCluster)
		wantErr bool
	}{
		{
			name: "Should create FlinkCluster with expected Metadata and Type",
			taskCtx: &FlinkTaskContext{
				ClusterName: "some-cluster-name",
				Namespace:   "some-namespace",
				Annotations: map[string]string{
					"annotation-1": "1",
					"annotation-2": "2",
				},
				Labels: map[string]string{
					"label-1": "1",
					"label-2": "2",
				},
			},
			check: func(t *testing.T, cluster *flinkOp.FlinkCluster) {
				assert.DeepEqual(t, cluster.ObjectMeta, metav1.ObjectMeta{
					Name:      "some-cluster-name",
					Namespace: "some-namespace",
					Annotations: map[string]string{
						"annotation-1": "1",
						"annotation-2": "2",
					},
					Labels: map[string]string{
						"label-1": "1",
						"label-2": "2",
					},
				})
				assert.DeepEqual(t, cluster.TypeMeta, metav1.TypeMeta{
					Kind:       "FlinkCluster",
					APIVersion: "flinkoperator.k8s.io/v1beta1",
				})
			},
		},
		{
			name: "Should create FlinkCluster with expected Spec",
			taskCtx: &FlinkTaskContext{
				ClusterName: "some-cluster-name",
				Namespace:   "some-namespace",
				Annotations: map[string]string{
					"annotation-1": "1",
					"annotation-2": "2",
				},
				Labels: map[string]string{
					"label-1":      "1",
					"execution-id": "1",
				},
				Job: flinkIdl.FlinkJob{
					MainClass:   "SomeClass",
					JarFiles:    exampleArtifacts,
					Parallelism: int32(10),
					FlinkProperties: map[string]string{
						"taskmanager.numberOfTaskSlots":            "1",
						"metrics.reporter.promgateway.groupingKey": `namespace={{.Namespace}};cluster={{.ClusterName}};execution_id={{index .Labels "execution-id"}}`,
					},
					Image: "flink-image",
					JobManager: &flinkIdl.JobManager{
						Resource: &flinkIdl.Resource{
							Cpu: &flinkIdl.Resource_Quantity{
								String_: "4",
							},
							Memory: &flinkIdl.Resource_Quantity{
								String_: "10Gi",
							},
							PersistentVolume: &flinkIdl.Resource_PersistentVolume{
								Type: flinkIdl.Resource_PersistentVolume_PD_SSD,
								Size: &flinkIdl.Resource_Quantity{String_: "100Gi"},
							},
						},
					},
					TaskManager: &flinkIdl.TaskManager{
						Resource: &flinkIdl.Resource{
							Cpu: &flinkIdl.Resource_Quantity{
								String_: "2",
							},
							Memory: &flinkIdl.Resource_Quantity{
								String_: "4Gi",
							},
							PersistentVolume: &flinkIdl.Resource_PersistentVolume{
								Type: flinkIdl.Resource_PersistentVolume_PD_SSD,
								Size: &flinkIdl.Resource_Quantity{String_: "20Gi"},
							},
						},
					},
					Jflyte: &flinkIdl.JFlyte{
						IndexFileLocation: "gs://bucket/index-file.json",
						Artifacts: []*flinkIdl.JFlyte_Artifact{
							{
								Location: "gs://bucket/artifact0.jar",
								Name:     "artifact0.jar",
							},
							{
								Location: "gs://bucket/artifact1.jar",
								Name:     "artifact1.jar",
							},
							{
								Location: "gs://bucket/artifact2.jar",
								Name:     "artifact2.jar",
							},
						},
					},
				},
			},
			check: func(t *testing.T, cluster *flinkOp.FlinkCluster) {
				expectedReplicas := int32(1)
				expectedParralelism := int32(10)
				className := "SomeClass"
				recreateOnUpdate := true
				storageClassName := "pd-ssd"

				assert.DeepEqual(t, cluster.Spec.Image, flinkOp.ImageSpec{Name: "flink-image", PullPolicy: "Always"})
				assert.DeepEqual(t, cluster.Spec.FlinkProperties, map[string]string{
					"io.tmp.dirs": "/flink-tmp",
					"metrics.reporter.promgateway.groupingKey": "namespace=some-namespace;cluster=some-cluster-name;execution_id=1",
					"taskmanager.numberOfTaskSlots":            "1",
				})
				assert.DeepEqual(t, cluster.Spec.RecreateOnUpdate, &recreateOnUpdate)
				assert.DeepEqual(t, cluster.Spec.EnvVars, []corev1.EnvVar{{Name: "STAGED_JARS", Value: "gs://bucket/a.jar gs://bucket/b.jar gs://bucket/c.jar"}})

				assert.DeepEqual(t, cluster.Spec.Job, &flinkOp.JobSpec{
					ClassName:      &className,
					Parallelism:    &expectedParralelism,
					InitContainers: []corev1.Container{},
					PodAnnotations: map[string]string{"annotation-1": "1", "annotation-2": "2"},
					PodLabels:      map[string]string{"execution-id": "1", "label-1": "1"},
				})

				assert.DeepEqual(t, cluster.Spec.JobManager, &flinkOp.JobManagerSpec{
					Replicas: &expectedReplicas,
					Resources: corev1.ResourceRequirements{Limits: corev1.ResourceList{
						"cpu":    resource.MustParse("4"),
						"memory": resource.MustParse("10Gi"),
					}},
					VolumeMounts: []corev1.VolumeMount{{Name: "pvc-jm", MountPath: "/flink-tmp"}},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
						{
							ObjectMeta: metav1.ObjectMeta{Name: "pvc-jm"},
							Spec: corev1.PersistentVolumeClaimSpec{
								AccessModes: []corev1.PersistentVolumeAccessMode{"ReadWriteOnce"},
								Resources: corev1.ResourceRequirements{Requests: corev1.ResourceList{
									"storage": resource.MustParse("100Gi"),
								}},
								StorageClassName: &storageClassName,
							},
						},
					},
					PodAnnotations: map[string]string{"annotation-1": "1", "annotation-2": "2"},
					PodLabels:      map[string]string{"execution-id": "1", "label-1": "1"},
				})

				assert.DeepEqual(t, cluster.Spec.TaskManager, &flinkOp.TaskManagerSpec{
					Resources: corev1.ResourceRequirements{Limits: corev1.ResourceList{
						"cpu":    resource.MustParse("2"),
						"memory": resource.MustParse("4Gi"),
					}},
					VolumeMounts: []corev1.VolumeMount{{Name: "pvc-tm", MountPath: "/flink-tmp"}},
					VolumeClaimTemplates: []corev1.PersistentVolumeClaim{
						{
							ObjectMeta: metav1.ObjectMeta{Name: "pvc-tm"},
							Spec: corev1.PersistentVolumeClaimSpec{
								AccessModes: []corev1.PersistentVolumeAccessMode{"ReadWriteOnce"},
								Resources: corev1.ResourceRequirements{Requests: corev1.ResourceList{
									"storage": resource.MustParse("20Gi"),
								}},
								StorageClassName: &storageClassName,
							},
						},
					},
					PodAnnotations: map[string]string{"annotation-1": "1", "annotation-2": "2"},
					PodLabels:      map[string]string{"execution-id": "1", "label-1": "1"},
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFlinkCluster(tt.taskCtx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFlinkCluster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.check(t, got)
		})
	}
}
