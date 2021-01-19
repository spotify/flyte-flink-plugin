.. _api_file_flyteidl-flink/flink.proto:

flink.proto
==========================

.. _api_msg_flyteidl_flink.JobManager:

flyteidl_flink.JobManager
-------------------------

`[flyteidl_flink.JobManager proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L8>`_


.. code-block:: json

  {
    "cpu": "{...}",
    "memory": "{...}"
  }

.. _api_field_flyteidl_flink.JobManager.cpu:

cpu
  (:ref:`k8s.io.apimachinery.pkg.api.resource.Quantity <api_msg_k8s.io.apimachinery.pkg.api.resource.Quantity>`) 
  
.. _api_field_flyteidl_flink.JobManager.memory:

memory
  (:ref:`k8s.io.apimachinery.pkg.api.resource.Quantity <api_msg_k8s.io.apimachinery.pkg.api.resource.Quantity>`) 
  


.. _api_msg_flyteidl_flink.TaskManager:

flyteidl_flink.TaskManager
--------------------------

`[flyteidl_flink.TaskManager proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L13>`_


.. code-block:: json

  {
    "cpu": "{...}",
    "memory": "{...}",
    "replicas": "..."
  }

.. _api_field_flyteidl_flink.TaskManager.cpu:

cpu
  (:ref:`k8s.io.apimachinery.pkg.api.resource.Quantity <api_msg_k8s.io.apimachinery.pkg.api.resource.Quantity>`) 
  
.. _api_field_flyteidl_flink.TaskManager.memory:

memory
  (:ref:`k8s.io.apimachinery.pkg.api.resource.Quantity <api_msg_k8s.io.apimachinery.pkg.api.resource.Quantity>`) 
  
.. _api_field_flyteidl_flink.TaskManager.replicas:

replicas
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) 
  


.. _api_msg_flyteidl_flink.FlinkJob:

flyteidl_flink.FlinkJob
-----------------------

`[flyteidl_flink.FlinkJob proto] <https://github.com/lyft/flyteidl/blob/master/protos/flyteidl-flink/flink.proto#L20>`_

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
  

