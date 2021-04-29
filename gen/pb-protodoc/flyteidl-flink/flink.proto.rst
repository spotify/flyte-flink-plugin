.. _api_file_flyteidl-flink/flink.proto:

flink.proto
==========================

.. _api_msg_flyteidl_flink.Resource:

flyteidl_flink.Resource
-----------------------

`[flyteidl_flink.Resource proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L8>`_


.. code-block:: json

  {
    "cpu": "{...}",
    "memory": "{...}",
    "persistentVolume": "{...}"
  }

.. _api_field_flyteidl_flink.Resource.cpu:

cpu
  (:ref:`flyteidl_flink.Resource.Quantity <api_msg_flyteidl_flink.Resource.Quantity>`) 
  
.. _api_field_flyteidl_flink.Resource.memory:

memory
  (:ref:`flyteidl_flink.Resource.Quantity <api_msg_flyteidl_flink.Resource.Quantity>`) 
  
.. _api_field_flyteidl_flink.Resource.persistentVolume:

persistentVolume
  (:ref:`flyteidl_flink.Resource.PersistentVolume <api_msg_flyteidl_flink.Resource.PersistentVolume>`) 
  
.. _api_msg_flyteidl_flink.Resource.Quantity:

flyteidl_flink.Resource.Quantity
--------------------------------

`[flyteidl_flink.Resource.Quantity proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L11>`_

Value must be a valid k8s quantity. See
https://github.com/kubernetes/apimachinery/blob/master/pkg/api/resource/quantity.go#L30-L80

.. code-block:: json

  {
    "string": "..."
  }

.. _api_field_flyteidl_flink.Resource.Quantity.string:

string
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  


.. _api_msg_flyteidl_flink.Resource.PersistentVolume:

flyteidl_flink.Resource.PersistentVolume
----------------------------------------

`[flyteidl_flink.Resource.PersistentVolume proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L17>`_


.. code-block:: json

  {
    "type": "...",
    "size": "{...}"
  }

.. _api_field_flyteidl_flink.Resource.PersistentVolume.type:

type
  (:ref:`flyteidl_flink.Resource.PersistentVolume.Type <api_enum_flyteidl_flink.Resource.PersistentVolume.Type>`) 
  
.. _api_field_flyteidl_flink.Resource.PersistentVolume.size:

size
  (:ref:`flyteidl_flink.Resource.Quantity <api_msg_flyteidl_flink.Resource.Quantity>`) 
  

.. _api_enum_flyteidl_flink.Resource.PersistentVolume.Type:

Enum flyteidl_flink.Resource.PersistentVolume.Type
--------------------------------------------------

`[flyteidl_flink.Resource.PersistentVolume.Type proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L18>`_


.. _api_enum_value_flyteidl_flink.Resource.PersistentVolume.Type.PD_STANDARD:

PD_STANDARD
  *(DEFAULT)* ⁣
  
.. _api_enum_value_flyteidl_flink.Resource.PersistentVolume.Type.PD_SSD:

PD_SSD
  ⁣
  


.. _api_msg_flyteidl_flink.JobManager:

flyteidl_flink.JobManager
-------------------------

`[flyteidl_flink.JobManager proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L32>`_


.. code-block:: json

  {
    "resource": "{...}"
  }

.. _api_field_flyteidl_flink.JobManager.resource:

resource
  (:ref:`flyteidl_flink.Resource <api_msg_flyteidl_flink.Resource>`) 
  


.. _api_msg_flyteidl_flink.TaskManager:

flyteidl_flink.TaskManager
--------------------------

`[flyteidl_flink.TaskManager proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L34>`_


.. code-block:: json

  {
    "resource": "{...}",
    "replicas": "..."
  }

.. _api_field_flyteidl_flink.TaskManager.resource:

resource
  (:ref:`flyteidl_flink.Resource <api_msg_flyteidl_flink.Resource>`) 
  
.. _api_field_flyteidl_flink.TaskManager.replicas:

replicas
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  


.. _api_msg_flyteidl_flink.JFlyte:

flyteidl_flink.JFlyte
---------------------

`[flyteidl_flink.JFlyte proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L39>`_


.. code-block:: json

  {
    "index_file_location": "...",
    "artifacts": []
  }

.. _api_field_flyteidl_flink.JFlyte.index_file_location:

index_file_location
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  
.. _api_field_flyteidl_flink.JFlyte.artifacts:

artifacts
  (:ref:`flyteidl_flink.JFlyte.Artifact <api_msg_flyteidl_flink.JFlyte.Artifact>`) 
  
.. _api_msg_flyteidl_flink.JFlyte.Artifact:

flyteidl_flink.JFlyte.Artifact
------------------------------

`[flyteidl_flink.JFlyte.Artifact proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L40>`_


.. code-block:: json

  {
    "name": "...",
    "location": "..."
  }

.. _api_field_flyteidl_flink.JFlyte.Artifact.name:

name
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  
.. _api_field_flyteidl_flink.JFlyte.Artifact.location:

location
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  



.. _api_msg_flyteidl_flink.FlinkJob:

flyteidl_flink.FlinkJob
-----------------------

`[flyteidl_flink.FlinkJob proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L51>`_

Custom Proto for Flink Plugin.

.. code-block:: json

  {
    "jarFiles": [],
    "mainClass": "...",
    "args": [],
    "flinkProperties": "{...}",
    "jobManager": "{...}",
    "taskManager": "{...}",
    "serviceAccount": "...",
    "image": "...",
    "jflyte": "{...}"
  }

.. _api_field_flyteidl_flink.FlinkJob.jarFiles:

jarFiles
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  
.. _api_field_flyteidl_flink.FlinkJob.mainClass:

mainClass
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  
.. _api_field_flyteidl_flink.FlinkJob.args:

args
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  
.. _api_field_flyteidl_flink.FlinkJob.flinkProperties:

flinkProperties
  (map<`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_, `string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_>) 
  
.. _api_field_flyteidl_flink.FlinkJob.jobManager:

jobManager
  (:ref:`flyteidl_flink.JobManager <api_msg_flyteidl_flink.JobManager>`) 
  
.. _api_field_flyteidl_flink.FlinkJob.taskManager:

taskManager
  (:ref:`flyteidl_flink.TaskManager <api_msg_flyteidl_flink.TaskManager>`) 
  
.. _api_field_flyteidl_flink.FlinkJob.serviceAccount:

serviceAccount
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  
.. _api_field_flyteidl_flink.FlinkJob.image:

image
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  
.. _api_field_flyteidl_flink.FlinkJob.jflyte:

jflyte
  (:ref:`flyteidl_flink.JFlyte <api_msg_flyteidl_flink.JFlyte>`) if using experiment flytekit-java this will contain all artifacts required
  to run the task
  
  


.. _api_msg_flyteidl_flink.JobExecutionInfo:

flyteidl_flink.JobExecutionInfo
-------------------------------

`[flyteidl_flink.JobExecutionInfo proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L70>`_


.. code-block:: json

  {
    "id": "..."
  }

.. _api_field_flyteidl_flink.JobExecutionInfo.id:

id
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  


.. _api_msg_flyteidl_flink.JobManagerExecutionInfo:

flyteidl_flink.JobManagerExecutionInfo
--------------------------------------

`[flyteidl_flink.JobManagerExecutionInfo proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L72>`_


.. code-block:: json

  {
    "ingressURLs": []
  }

.. _api_field_flyteidl_flink.JobManagerExecutionInfo.ingressURLs:

ingressURLs
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  


.. _api_msg_flyteidl_flink.FlinkExecutionInfo:

flyteidl_flink.FlinkExecutionInfo
---------------------------------

`[flyteidl_flink.FlinkExecutionInfo proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L74>`_


.. code-block:: json

  {
    "job": "{...}",
    "jobManager": "{...}"
  }

.. _api_field_flyteidl_flink.FlinkExecutionInfo.job:

job
  (:ref:`flyteidl_flink.JobExecutionInfo <api_msg_flyteidl_flink.JobExecutionInfo>`) 
  
.. _api_field_flyteidl_flink.FlinkExecutionInfo.jobManager:

jobManager
  (:ref:`flyteidl_flink.JobManagerExecutionInfo <api_msg_flyteidl_flink.JobManagerExecutionInfo>`) 
  

