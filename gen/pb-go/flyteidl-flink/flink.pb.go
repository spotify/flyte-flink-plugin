// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl-flink/flink.proto

package flyteidl_flink

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Resource_PersistentVolume_Type int32

const (
	Resource_PersistentVolume_PD_STANDARD Resource_PersistentVolume_Type = 0
	Resource_PersistentVolume_PD_SSD      Resource_PersistentVolume_Type = 1
)

var Resource_PersistentVolume_Type_name = map[int32]string{
	0: "PD_STANDARD",
	1: "PD_SSD",
}

var Resource_PersistentVolume_Type_value = map[string]int32{
	"PD_STANDARD": 0,
	"PD_SSD":      1,
}

func (x Resource_PersistentVolume_Type) String() string {
	return proto.EnumName(Resource_PersistentVolume_Type_name, int32(x))
}

func (Resource_PersistentVolume_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{0, 1, 0}
}

type Resource struct {
	Cpu                  *Resource_Quantity         `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory               *Resource_Quantity         `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	PersistentVolume     *Resource_PersistentVolume `protobuf:"bytes,3,opt,name=persistentVolume,proto3" json:"persistentVolume,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetCpu() *Resource_Quantity {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Resource) GetMemory() *Resource_Quantity {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *Resource) GetPersistentVolume() *Resource_PersistentVolume {
	if m != nil {
		return m.PersistentVolume
	}
	return nil
}

// Value must be a valid k8s quantity. See
// https://github.com/kubernetes/apimachinery/blob/master/pkg/api/resource/quantity.go#L30-L80
type Resource_Quantity struct {
	String_              string   `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource_Quantity) Reset()         { *m = Resource_Quantity{} }
func (m *Resource_Quantity) String() string { return proto.CompactTextString(m) }
func (*Resource_Quantity) ProtoMessage()    {}
func (*Resource_Quantity) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{0, 0}
}

func (m *Resource_Quantity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource_Quantity.Unmarshal(m, b)
}
func (m *Resource_Quantity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource_Quantity.Marshal(b, m, deterministic)
}
func (m *Resource_Quantity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource_Quantity.Merge(m, src)
}
func (m *Resource_Quantity) XXX_Size() int {
	return xxx_messageInfo_Resource_Quantity.Size(m)
}
func (m *Resource_Quantity) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource_Quantity.DiscardUnknown(m)
}

var xxx_messageInfo_Resource_Quantity proto.InternalMessageInfo

func (m *Resource_Quantity) GetString_() string {
	if m != nil {
		return m.String_
	}
	return ""
}

type Resource_PersistentVolume struct {
	Type                 Resource_PersistentVolume_Type `protobuf:"varint,1,opt,name=type,proto3,enum=flyteidl_flink.Resource_PersistentVolume_Type" json:"type,omitempty"`
	Size                 *Resource_Quantity             `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Resource_PersistentVolume) Reset()         { *m = Resource_PersistentVolume{} }
func (m *Resource_PersistentVolume) String() string { return proto.CompactTextString(m) }
func (*Resource_PersistentVolume) ProtoMessage()    {}
func (*Resource_PersistentVolume) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{0, 1}
}

func (m *Resource_PersistentVolume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource_PersistentVolume.Unmarshal(m, b)
}
func (m *Resource_PersistentVolume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource_PersistentVolume.Marshal(b, m, deterministic)
}
func (m *Resource_PersistentVolume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource_PersistentVolume.Merge(m, src)
}
func (m *Resource_PersistentVolume) XXX_Size() int {
	return xxx_messageInfo_Resource_PersistentVolume.Size(m)
}
func (m *Resource_PersistentVolume) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource_PersistentVolume.DiscardUnknown(m)
}

var xxx_messageInfo_Resource_PersistentVolume proto.InternalMessageInfo

func (m *Resource_PersistentVolume) GetType() Resource_PersistentVolume_Type {
	if m != nil {
		return m.Type
	}
	return Resource_PersistentVolume_PD_STANDARD
}

func (m *Resource_PersistentVolume) GetSize() *Resource_Quantity {
	if m != nil {
		return m.Size
	}
	return nil
}

type JobManager struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *JobManager) Reset()         { *m = JobManager{} }
func (m *JobManager) String() string { return proto.CompactTextString(m) }
func (*JobManager) ProtoMessage()    {}
func (*JobManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{1}
}

func (m *JobManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobManager.Unmarshal(m, b)
}
func (m *JobManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobManager.Marshal(b, m, deterministic)
}
func (m *JobManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobManager.Merge(m, src)
}
func (m *JobManager) XXX_Size() int {
	return xxx_messageInfo_JobManager.Size(m)
}
func (m *JobManager) XXX_DiscardUnknown() {
	xxx_messageInfo_JobManager.DiscardUnknown(m)
}

var xxx_messageInfo_JobManager proto.InternalMessageInfo

func (m *JobManager) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type TaskManager struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Replicas             int32     `protobuf:"varint,2,opt,name=replicas,proto3" json:"replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TaskManager) Reset()         { *m = TaskManager{} }
func (m *TaskManager) String() string { return proto.CompactTextString(m) }
func (*TaskManager) ProtoMessage()    {}
func (*TaskManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{2}
}

func (m *TaskManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskManager.Unmarshal(m, b)
}
func (m *TaskManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskManager.Marshal(b, m, deterministic)
}
func (m *TaskManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskManager.Merge(m, src)
}
func (m *TaskManager) XXX_Size() int {
	return xxx_messageInfo_TaskManager.Size(m)
}
func (m *TaskManager) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskManager.DiscardUnknown(m)
}

var xxx_messageInfo_TaskManager proto.InternalMessageInfo

func (m *TaskManager) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *TaskManager) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

type JFlyte struct {
	IndexFileLocation    string             `protobuf:"bytes,1,opt,name=index_file_location,json=indexFileLocation,proto3" json:"index_file_location,omitempty"`
	Artifacts            []*JFlyte_Artifact `protobuf:"bytes,2,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *JFlyte) Reset()         { *m = JFlyte{} }
func (m *JFlyte) String() string { return proto.CompactTextString(m) }
func (*JFlyte) ProtoMessage()    {}
func (*JFlyte) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{3}
}

func (m *JFlyte) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JFlyte.Unmarshal(m, b)
}
func (m *JFlyte) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JFlyte.Marshal(b, m, deterministic)
}
func (m *JFlyte) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JFlyte.Merge(m, src)
}
func (m *JFlyte) XXX_Size() int {
	return xxx_messageInfo_JFlyte.Size(m)
}
func (m *JFlyte) XXX_DiscardUnknown() {
	xxx_messageInfo_JFlyte.DiscardUnknown(m)
}

var xxx_messageInfo_JFlyte proto.InternalMessageInfo

func (m *JFlyte) GetIndexFileLocation() string {
	if m != nil {
		return m.IndexFileLocation
	}
	return ""
}

func (m *JFlyte) GetArtifacts() []*JFlyte_Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type JFlyte_Artifact struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location             string   `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JFlyte_Artifact) Reset()         { *m = JFlyte_Artifact{} }
func (m *JFlyte_Artifact) String() string { return proto.CompactTextString(m) }
func (*JFlyte_Artifact) ProtoMessage()    {}
func (*JFlyte_Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{3, 0}
}

func (m *JFlyte_Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JFlyte_Artifact.Unmarshal(m, b)
}
func (m *JFlyte_Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JFlyte_Artifact.Marshal(b, m, deterministic)
}
func (m *JFlyte_Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JFlyte_Artifact.Merge(m, src)
}
func (m *JFlyte_Artifact) XXX_Size() int {
	return xxx_messageInfo_JFlyte_Artifact.Size(m)
}
func (m *JFlyte_Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_JFlyte_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_JFlyte_Artifact proto.InternalMessageInfo

func (m *JFlyte_Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JFlyte_Artifact) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

// Custom Proto for Flink Plugin.
type FlinkJob struct {
	JarFiles                []string          `protobuf:"bytes,1,rep,name=jarFiles,proto3" json:"jarFiles,omitempty"`
	MainClass               string            `protobuf:"bytes,2,opt,name=mainClass,proto3" json:"mainClass,omitempty"`
	Args                    []string          `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	FlinkProperties         map[string]string `protobuf:"bytes,4,rep,name=flinkProperties,proto3" json:"flinkProperties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	JobManager              *JobManager       `protobuf:"bytes,5,opt,name=jobManager,proto3" json:"jobManager,omitempty"`
	TaskManager             *TaskManager      `protobuf:"bytes,6,opt,name=taskManager,proto3" json:"taskManager,omitempty"`
	ServiceAccount          string            `protobuf:"bytes,7,opt,name=serviceAccount,proto3" json:"serviceAccount,omitempty"`
	Image                   string            `protobuf:"bytes,8,opt,name=image,proto3" json:"image,omitempty"`
	FlinkVersion            string            `protobuf:"bytes,9,opt,name=flinkVersion,proto3" json:"flinkVersion,omitempty"`
	Parallelism             int32             `protobuf:"varint,10,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	KubernetesClusterName   string            `protobuf:"bytes,11,opt,name=kubernetesClusterName,proto3" json:"kubernetesClusterName,omitempty"`
	KubernetesClusterRegion string            `protobuf:"bytes,12,opt,name=kubernetesClusterRegion,proto3" json:"kubernetesClusterRegion,omitempty"`
	// if using experiment flytekit-java this will contain all artifacts required
	// to run the task
	Jflyte               *JFlyte  `protobuf:"bytes,100,opt,name=jflyte,proto3" json:"jflyte,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlinkJob) Reset()         { *m = FlinkJob{} }
func (m *FlinkJob) String() string { return proto.CompactTextString(m) }
func (*FlinkJob) ProtoMessage()    {}
func (*FlinkJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{4}
}

func (m *FlinkJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlinkJob.Unmarshal(m, b)
}
func (m *FlinkJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlinkJob.Marshal(b, m, deterministic)
}
func (m *FlinkJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlinkJob.Merge(m, src)
}
func (m *FlinkJob) XXX_Size() int {
	return xxx_messageInfo_FlinkJob.Size(m)
}
func (m *FlinkJob) XXX_DiscardUnknown() {
	xxx_messageInfo_FlinkJob.DiscardUnknown(m)
}

var xxx_messageInfo_FlinkJob proto.InternalMessageInfo

func (m *FlinkJob) GetJarFiles() []string {
	if m != nil {
		return m.JarFiles
	}
	return nil
}

func (m *FlinkJob) GetMainClass() string {
	if m != nil {
		return m.MainClass
	}
	return ""
}

func (m *FlinkJob) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *FlinkJob) GetFlinkProperties() map[string]string {
	if m != nil {
		return m.FlinkProperties
	}
	return nil
}

func (m *FlinkJob) GetJobManager() *JobManager {
	if m != nil {
		return m.JobManager
	}
	return nil
}

func (m *FlinkJob) GetTaskManager() *TaskManager {
	if m != nil {
		return m.TaskManager
	}
	return nil
}

func (m *FlinkJob) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *FlinkJob) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *FlinkJob) GetFlinkVersion() string {
	if m != nil {
		return m.FlinkVersion
	}
	return ""
}

func (m *FlinkJob) GetParallelism() int32 {
	if m != nil {
		return m.Parallelism
	}
	return 0
}

func (m *FlinkJob) GetKubernetesClusterName() string {
	if m != nil {
		return m.KubernetesClusterName
	}
	return ""
}

func (m *FlinkJob) GetKubernetesClusterRegion() string {
	if m != nil {
		return m.KubernetesClusterRegion
	}
	return ""
}

func (m *FlinkJob) GetJflyte() *JFlyte {
	if m != nil {
		return m.Jflyte
	}
	return nil
}

type JobExecutionInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobExecutionInfo) Reset()         { *m = JobExecutionInfo{} }
func (m *JobExecutionInfo) String() string { return proto.CompactTextString(m) }
func (*JobExecutionInfo) ProtoMessage()    {}
func (*JobExecutionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{5}
}

func (m *JobExecutionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobExecutionInfo.Unmarshal(m, b)
}
func (m *JobExecutionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobExecutionInfo.Marshal(b, m, deterministic)
}
func (m *JobExecutionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobExecutionInfo.Merge(m, src)
}
func (m *JobExecutionInfo) XXX_Size() int {
	return xxx_messageInfo_JobExecutionInfo.Size(m)
}
func (m *JobExecutionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JobExecutionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JobExecutionInfo proto.InternalMessageInfo

func (m *JobExecutionInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type JobManagerExecutionInfo struct {
	IngressURLs          []string `protobuf:"bytes,1,rep,name=ingressURLs,proto3" json:"ingressURLs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobManagerExecutionInfo) Reset()         { *m = JobManagerExecutionInfo{} }
func (m *JobManagerExecutionInfo) String() string { return proto.CompactTextString(m) }
func (*JobManagerExecutionInfo) ProtoMessage()    {}
func (*JobManagerExecutionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{6}
}

func (m *JobManagerExecutionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobManagerExecutionInfo.Unmarshal(m, b)
}
func (m *JobManagerExecutionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobManagerExecutionInfo.Marshal(b, m, deterministic)
}
func (m *JobManagerExecutionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobManagerExecutionInfo.Merge(m, src)
}
func (m *JobManagerExecutionInfo) XXX_Size() int {
	return xxx_messageInfo_JobManagerExecutionInfo.Size(m)
}
func (m *JobManagerExecutionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JobManagerExecutionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JobManagerExecutionInfo proto.InternalMessageInfo

func (m *JobManagerExecutionInfo) GetIngressURLs() []string {
	if m != nil {
		return m.IngressURLs
	}
	return nil
}

type FlinkExecutionInfo struct {
	Job                  *JobExecutionInfo        `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	JobManager           *JobManagerExecutionInfo `protobuf:"bytes,2,opt,name=jobManager,proto3" json:"jobManager,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FlinkExecutionInfo) Reset()         { *m = FlinkExecutionInfo{} }
func (m *FlinkExecutionInfo) String() string { return proto.CompactTextString(m) }
func (*FlinkExecutionInfo) ProtoMessage()    {}
func (*FlinkExecutionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f39d29671649d9, []int{7}
}

func (m *FlinkExecutionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlinkExecutionInfo.Unmarshal(m, b)
}
func (m *FlinkExecutionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlinkExecutionInfo.Marshal(b, m, deterministic)
}
func (m *FlinkExecutionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlinkExecutionInfo.Merge(m, src)
}
func (m *FlinkExecutionInfo) XXX_Size() int {
	return xxx_messageInfo_FlinkExecutionInfo.Size(m)
}
func (m *FlinkExecutionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FlinkExecutionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FlinkExecutionInfo proto.InternalMessageInfo

func (m *FlinkExecutionInfo) GetJob() *JobExecutionInfo {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *FlinkExecutionInfo) GetJobManager() *JobManagerExecutionInfo {
	if m != nil {
		return m.JobManager
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl_flink.Resource_PersistentVolume_Type", Resource_PersistentVolume_Type_name, Resource_PersistentVolume_Type_value)
	proto.RegisterType((*Resource)(nil), "flyteidl_flink.Resource")
	proto.RegisterType((*Resource_Quantity)(nil), "flyteidl_flink.Resource.Quantity")
	proto.RegisterType((*Resource_PersistentVolume)(nil), "flyteidl_flink.Resource.PersistentVolume")
	proto.RegisterType((*JobManager)(nil), "flyteidl_flink.JobManager")
	proto.RegisterType((*TaskManager)(nil), "flyteidl_flink.TaskManager")
	proto.RegisterType((*JFlyte)(nil), "flyteidl_flink.JFlyte")
	proto.RegisterType((*JFlyte_Artifact)(nil), "flyteidl_flink.JFlyte.Artifact")
	proto.RegisterType((*FlinkJob)(nil), "flyteidl_flink.FlinkJob")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl_flink.FlinkJob.FlinkPropertiesEntry")
	proto.RegisterType((*JobExecutionInfo)(nil), "flyteidl_flink.JobExecutionInfo")
	proto.RegisterType((*JobManagerExecutionInfo)(nil), "flyteidl_flink.JobManagerExecutionInfo")
	proto.RegisterType((*FlinkExecutionInfo)(nil), "flyteidl_flink.FlinkExecutionInfo")
}

func init() { proto.RegisterFile("flyteidl-flink/flink.proto", fileDescriptor_13f39d29671649d9) }

var fileDescriptor_13f39d29671649d9 = []byte{
	// 874 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xae, 0xfc, 0x17, 0xf9, 0x38, 0x70, 0x35, 0xae, 0x5b, 0x04, 0xef, 0x62, 0x99, 0xba, 0x75,
	0x4e, 0x33, 0xab, 0x85, 0xd3, 0x01, 0x6d, 0x87, 0x62, 0x88, 0x9a, 0xa4, 0x58, 0x9a, 0x16, 0x19,
	0x9b, 0xb6, 0xc0, 0x82, 0x2c, 0xa0, 0x65, 0xda, 0x63, 0x2c, 0x53, 0x06, 0x49, 0x05, 0xf5, 0x2e,
	0xf7, 0x04, 0x5b, 0x9f, 0x67, 0x40, 0x9e, 0x67, 0xaf, 0xe0, 0xab, 0x81, 0x94, 0xfc, 0x23, 0x27,
	0x01, 0x02, 0xf4, 0xc6, 0xa0, 0xce, 0xf7, 0x9d, 0xef, 0x1c, 0x9d, 0xf3, 0x99, 0x82, 0x46, 0x2f,
	0x1a, 0x2b, 0xca, 0xba, 0x51, 0xab, 0x17, 0x31, 0x3e, 0x78, 0x60, 0x7e, 0xfd, 0x91, 0x88, 0x55,
	0x8c, 0xea, 0x53, 0xec, 0xd4, 0x44, 0x1b, 0x6b, 0xe7, 0x24, 0x62, 0x5d, 0xa2, 0xe8, 0x83, 0xe9,
	0x21, 0x25, 0x7a, 0xff, 0x15, 0xc1, 0xc6, 0x54, 0xc6, 0x89, 0x08, 0x29, 0xda, 0x82, 0x62, 0x38,
	0x4a, 0x5c, 0x6b, 0xdd, 0x6a, 0xd6, 0xda, 0xdf, 0xf8, 0x79, 0x0d, 0x7f, 0x4a, 0xf3, 0x7f, 0x4d,
	0x08, 0x57, 0x4c, 0x8d, 0xb1, 0x66, 0xa3, 0x27, 0x50, 0x19, 0xd2, 0x61, 0x2c, 0xc6, 0x6e, 0xe1,
	0xa6, 0x79, 0x59, 0x02, 0x7a, 0x0b, 0xce, 0x88, 0x0a, 0xc9, 0xa4, 0xa2, 0x5c, 0xbd, 0x8b, 0xa3,
	0x64, 0x48, 0xdd, 0xa2, 0x11, 0xd9, 0xb8, 0x56, 0xe4, 0x70, 0x29, 0x01, 0x5f, 0x92, 0x68, 0xbc,
	0x07, 0x7b, 0x5a, 0x0a, 0xbd, 0x84, 0x8a, 0x54, 0x82, 0xf1, 0xbe, 0x79, 0xab, 0x6a, 0xb0, 0x35,
	0x09, 0x1e, 0x0a, 0xbf, 0xfd, 0xc3, 0xef, 0xcd, 0xe3, 0xcd, 0xd6, 0xc9, 0xcf, 0xc7, 0x0f, 0x5b,
	0x4f, 0xfc, 0x93, 0xcd, 0x8d, 0xe6, 0x31, 0xdd, 0x65, 0x3c, 0x19, 0x0e, 0x5e, 0xbe, 0x7a, 0x71,
	0x74, 0x78, 0x72, 0xff, 0xb8, 0xb5, 0x99, 0x82, 0x27, 0xf7, 0x37, 0xbe, 0xc5, 0x99, 0x44, 0xe3,
	0x5f, 0x0b, 0x9c, 0xe5, 0xfa, 0xe8, 0x00, 0x4a, 0x6a, 0x3c, 0xa2, 0x46, 0xbf, 0xde, 0xf6, 0x6f,
	0xdc, 0xb8, 0x7f, 0x34, 0x1e, 0xd1, 0xc0, 0x9e, 0x04, 0xe5, 0xbf, 0xac, 0x82, 0x63, 0x61, 0xa3,
	0x82, 0x7e, 0x84, 0x92, 0x64, 0x7f, 0xd2, 0x9b, 0xcf, 0xd2, 0xd0, 0xbd, 0xbb, 0x50, 0xd2, 0x72,
	0xe8, 0x36, 0xd4, 0x0e, 0x77, 0x4e, 0xdf, 0x1c, 0x6d, 0xbf, 0xde, 0xd9, 0xc6, 0x3b, 0xce, 0x2d,
	0x04, 0x50, 0xd1, 0x81, 0x37, 0x3b, 0x8e, 0xe5, 0x05, 0x00, 0xfb, 0x71, 0xe7, 0x15, 0xe1, 0xa4,
	0x4f, 0x05, 0x7a, 0x04, 0xb6, 0xc8, 0xd4, 0xb2, 0x8d, 0xbb, 0xd7, 0x55, 0xc3, 0x33, 0xa6, 0xf7,
	0x07, 0xd4, 0x8e, 0x88, 0x1c, 0x7c, 0x92, 0x08, 0xba, 0xab, 0xb3, 0x46, 0x11, 0x0b, 0x89, 0x34,
	0x2f, 0x5a, 0x0e, 0x56, 0x26, 0x41, 0xa9, 0x51, 0x68, 0xde, 0xc2, 0x33, 0xc0, 0xbb, 0xb0, 0xa0,
	0xb2, 0xbf, 0xa7, 0xb5, 0x90, 0x0f, 0x9f, 0x33, 0xde, 0xa5, 0x1f, 0x4e, 0x7b, 0x2c, 0xa2, 0xa7,
	0x51, 0x1c, 0x12, 0xc5, 0x62, 0x9e, 0x6e, 0x14, 0x7f, 0x66, 0xa0, 0x3d, 0x16, 0xd1, 0x83, 0x0c,
	0x40, 0xcf, 0xa0, 0x4a, 0x84, 0x62, 0x3d, 0x12, 0x2a, 0x5d, 0xa0, 0xd8, 0xac, 0xb5, 0xbf, 0x5e,
	0x6e, 0x2b, 0x95, 0xf6, 0xb7, 0x33, 0x1e, 0x9e, 0x67, 0x34, 0xf6, 0xc0, 0x9e, 0x86, 0x11, 0x82,
	0x12, 0x27, 0x43, 0x9a, 0xd5, 0x32, 0x67, 0x74, 0x0f, 0xec, 0x59, 0x0f, 0x05, 0xe3, 0x2a, 0x98,
	0x04, 0x2b, 0xa2, 0xec, 0x58, 0x7f, 0x5b, 0x16, 0x9e, 0x61, 0xde, 0x45, 0x19, 0xec, 0x3d, 0x5d,
	0x6c, 0x3f, 0xee, 0xa0, 0x2d, 0xb0, 0xcf, 0x88, 0xd0, 0x6d, 0x4a, 0xd7, 0x5a, 0x2f, 0x36, 0xab,
	0xc1, 0xda, 0x24, 0x28, 0x7f, 0xb4, 0x0a, 0xae, 0x35, 0x09, 0x56, 0x3f, 0x5a, 0x55, 0x6f, 0xae,
	0x30, 0x25, 0xa2, 0xef, 0xa0, 0x3a, 0x24, 0x8c, 0x3f, 0x8f, 0x88, 0x94, 0x59, 0x29, 0x3d, 0x29,
	0xa1, 0xfd, 0x32, 0x47, 0x74, 0x93, 0x44, 0xf4, 0xa5, 0x5b, 0xd4, 0xba, 0xd8, 0x9c, 0xd1, 0x7b,
	0xb8, 0x6d, 0x5e, 0xf4, 0x50, 0xc4, 0x23, 0x2a, 0x14, 0xa3, 0xd2, 0x2d, 0x99, 0x49, 0xb4, 0x96,
	0x27, 0x31, 0x6d, 0x31, 0x3d, 0xcc, 0xf9, 0xbb, 0x5c, 0x89, 0x31, 0x5e, 0x56, 0x41, 0x4f, 0x01,
	0xce, 0x66, 0x2e, 0x72, 0xcb, 0x66, 0xe9, 0x8d, 0x4b, 0xd3, 0x9d, 0x31, 0xf0, 0x02, 0x1b, 0x3d,
	0x83, 0x9a, 0x9a, 0xbb, 0xc7, 0xad, 0x98, 0xe4, 0xaf, 0x96, 0x93, 0x17, 0x0c, 0x86, 0x17, 0xf9,
	0xe8, 0x1e, 0xd4, 0x25, 0x15, 0xe7, 0x2c, 0xa4, 0xdb, 0x61, 0x18, 0x27, 0x5c, 0xb9, 0x2b, 0x66,
	0x2d, 0x4b, 0x51, 0x74, 0x07, 0xca, 0x6c, 0x48, 0xfa, 0xd4, 0xb5, 0x0d, 0x9c, 0x3e, 0x20, 0x0f,
	0x56, 0x8d, 0xfe, 0x3b, 0xfd, 0x47, 0x8c, 0xb9, 0x5b, 0x35, 0x60, 0x2e, 0x86, 0x36, 0xa0, 0x36,
	0x22, 0x82, 0x44, 0x11, 0x8d, 0x98, 0x1c, 0xba, 0x90, 0x37, 0xe7, 0x22, 0x86, 0x1e, 0xc1, 0x17,
	0x83, 0xa4, 0x43, 0x05, 0xa7, 0x8a, 0xca, 0xe7, 0x51, 0x22, 0x15, 0x15, 0xaf, 0xb5, 0x55, 0x6a,
	0x46, 0xf7, 0x6a, 0x10, 0x3d, 0x86, 0xb5, 0x4b, 0x00, 0xa6, 0x7d, 0xdd, 0xcf, 0xaa, 0xc9, 0xbb,
	0x0e, 0x46, 0x3e, 0x54, 0xce, 0xcc, 0xa0, 0xdc, 0xae, 0x19, 0xdb, 0x97, 0x57, 0x3b, 0x1a, 0x67,
	0xac, 0x46, 0x00, 0x77, 0xae, 0x5a, 0x28, 0x72, 0xa0, 0x38, 0xa0, 0xe3, 0xcc, 0xd0, 0xfa, 0xa8,
	0xc7, 0x75, 0x4e, 0xa2, 0x24, 0xbd, 0x74, 0xaa, 0x38, 0x7d, 0x78, 0x5a, 0x78, 0x6c, 0x79, 0x1e,
	0x38, 0xfb, 0x71, 0x67, 0xf7, 0x03, 0x0d, 0x13, 0xed, 0xe8, 0x5f, 0x78, 0x2f, 0x46, 0x75, 0x28,
	0xb0, 0x6e, 0x96, 0x5e, 0x60, 0x5d, 0xef, 0x27, 0x58, 0x9b, 0x6f, 0x3b, 0x4f, 0x5d, 0x87, 0x1a,
	0xe3, 0x7d, 0x41, 0xa5, 0x7c, 0x8b, 0x0f, 0x32, 0xdb, 0xe3, 0xc5, 0x90, 0xf7, 0x8f, 0x05, 0xc8,
	0x74, 0x99, 0x4f, 0x6c, 0x43, 0xf1, 0x2c, 0xee, 0x64, 0x37, 0xca, 0xfa, 0x15, 0xe6, 0xca, 0xd1,
	0xb1, 0x26, 0xa3, 0x17, 0x39, 0x5f, 0xa6, 0xf7, 0xe7, 0xf7, 0xd7, 0xfb, 0x32, 0xaf, 0xb0, 0x90,
	0x1a, 0x38, 0xbf, 0xd5, 0xf3, 0x5f, 0xd6, 0x4e, 0xc5, 0x7c, 0x2b, 0xb7, 0xfe, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0xa2, 0x5a, 0xf0, 0x29, 0x72, 0x07, 0x00, 0x00,
}
