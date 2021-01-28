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
  (:ref:`k8s.io.apimachinery.pkg.api.resource.Quantity <api_msg_k8s.io.apimachinery.pkg.api.resource.Quantity>`) 
  
.. _api_field_flyteidl_flink.Resource.memory:

memory
  (:ref:`k8s.io.apimachinery.pkg.api.resource.Quantity <api_msg_k8s.io.apimachinery.pkg.api.resource.Quantity>`) 
  
.. _api_field_flyteidl_flink.Resource.persistentVolume:

persistentVolume
  (:ref:`flyteidl_flink.Resource.PersistentVolume <api_msg_flyteidl_flink.Resource.PersistentVolume>`) 
  
.. _api_msg_flyteidl_flink.Resource.PersistentVolume:

flyteidl_flink.Resource.PersistentVolume
----------------------------------------

`[flyteidl_flink.Resource.PersistentVolume proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L9>`_


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
  (:ref:`k8s.io.apimachinery.pkg.api.resource.Quantity <api_msg_k8s.io.apimachinery.pkg.api.resource.Quantity>`) 
  

.. _api_enum_flyteidl_flink.Resource.PersistentVolume.Type:

Enum flyteidl_flink.Resource.PersistentVolume.Type
--------------------------------------------------

`[flyteidl_flink.Resource.PersistentVolume.Type proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L10>`_


.. _api_enum_value_flyteidl_flink.Resource.PersistentVolume.Type.PD_STANDARD:

PD_STANDARD
  *(DEFAULT)* ⁣
  
.. _api_enum_value_flyteidl_flink.Resource.PersistentVolume.Type.PD_SSD:

PD_SSD
  ⁣
  


.. _api_msg_flyteidl_flink.JobManager:

flyteidl_flink.JobManager
-------------------------

`[flyteidl_flink.JobManager proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L24>`_


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

`[flyteidl_flink.TaskManager proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L28>`_


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
  


.. _api_msg_flyteidl_flink.FlinkJob:

flyteidl_flink.FlinkJob
-----------------------

`[flyteidl_flink.FlinkJob proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L34>`_

Custom Proto for Flink Plugin.

.. code-block:: json

  {
    "jarFile": "...",
    "mainClass": "...",
    "args": [],
    "flinkProperties": "{...}",
    "jobManager": "{...}",
    "taskManager": "{...}"
  }

.. _api_field_flyteidl_flink.FlinkJob.jarFile:

jarFile
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
  

