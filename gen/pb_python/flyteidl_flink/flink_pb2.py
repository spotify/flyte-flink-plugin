# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl-flink/flink.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from k8s.io.apimachinery.pkg.api.resource import generated_pb2 as k8s_dot_io_dot_apimachinery_dot_pkg_dot_api_dot_resource_dot_generated__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl-flink/flink.proto',
  package='flyteidl_flink',
  syntax='proto3',
  serialized_options=_b('Z>github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flink'),
  serialized_pb=_b('\n\x1a\x66lyteidl-flink/flink.proto\x12\x0e\x66lyteidl_flink\x1a\x34k8s.io/apimachinery/pkg/api/resource/generated.proto\"\x82\x03\n\x08Resource\x12;\n\x03\x63pu\x18\x01 \x01(\x0b\x32..k8s.io.apimachinery.pkg.api.resource.Quantity\x12>\n\x06memory\x18\x02 \x01(\x0b\x32..k8s.io.apimachinery.pkg.api.resource.Quantity\x12\x43\n\x10persistentVolume\x18\x03 \x01(\x0b\x32).flyteidl_flink.Resource.PersistentVolume\x1a\xb3\x01\n\x10PersistentVolume\x12<\n\x04type\x18\x01 \x01(\x0e\x32..flyteidl_flink.Resource.PersistentVolume.Type\x12<\n\x04size\x18\x02 \x01(\x0b\x32..k8s.io.apimachinery.pkg.api.resource.Quantity\"#\n\x04Type\x12\x0f\n\x0bPD_STANDARD\x10\x00\x12\n\n\x06PD_SSD\x10\x01\"8\n\nJobManager\x12*\n\x08resource\x18\x01 \x01(\x0b\x32\x18.flyteidl_flink.Resource\"K\n\x0bTaskManager\x12*\n\x08resource\x18\x01 \x01(\x0b\x32\x18.flyteidl_flink.Resource\x12\x10\n\x08replicas\x18\x02 \x01(\x05\"\xc5\x02\n\x08\x46linkJob\x12\x0f\n\x07jarFile\x18\x01 \x01(\t\x12\x11\n\tmainClass\x18\x02 \x01(\t\x12\x0c\n\x04\x61rgs\x18\x03 \x03(\t\x12\x46\n\x0f\x66linkProperties\x18\x04 \x03(\x0b\x32-.flyteidl_flink.FlinkJob.FlinkPropertiesEntry\x12.\n\njobManager\x18\x05 \x01(\x0b\x32\x1a.flyteidl_flink.JobManager\x12\x30\n\x0btaskManager\x18\x06 \x01(\x0b\x32\x1b.flyteidl_flink.TaskManager\x12\x16\n\x0eserviceAccount\x18\x07 \x01(\t\x12\r\n\x05image\x18\x08 \x01(\t\x1a\x36\n\x14\x46linkPropertiesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x1e\n\x10JobExecutionInfo\x12\n\n\x02id\x18\x01 \x01(\t\".\n\x17JobManagerExecutionInfo\x12\x13\n\x0bingressURLs\x18\x01 \x03(\t\"\x80\x01\n\x12\x46linkExecutionInfo\x12-\n\x03job\x18\x01 \x01(\x0b\x32 .flyteidl_flink.JobExecutionInfo\x12;\n\njobManager\x18\x02 \x01(\x0b\x32\'.flyteidl_flink.JobManagerExecutionInfoB@Z>github.com/spotify/flyte-flink-plugin/gen/pb-go/flyteidl-flinkb\x06proto3')
  ,
  dependencies=[k8s_dot_io_dot_apimachinery_dot_pkg_dot_api_dot_resource_dot_generated__pb2.DESCRIPTOR,])



_RESOURCE_PERSISTENTVOLUME_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='flyteidl_flink.Resource.PersistentVolume.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PD_STANDARD', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PD_SSD', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=452,
  serialized_end=487,
)
_sym_db.RegisterEnumDescriptor(_RESOURCE_PERSISTENTVOLUME_TYPE)


_RESOURCE_PERSISTENTVOLUME = _descriptor.Descriptor(
  name='PersistentVolume',
  full_name='flyteidl_flink.Resource.PersistentVolume',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='flyteidl_flink.Resource.PersistentVolume.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='size', full_name='flyteidl_flink.Resource.PersistentVolume.size', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _RESOURCE_PERSISTENTVOLUME_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=308,
  serialized_end=487,
)

_RESOURCE = _descriptor.Descriptor(
  name='Resource',
  full_name='flyteidl_flink.Resource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cpu', full_name='flyteidl_flink.Resource.cpu', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='memory', full_name='flyteidl_flink.Resource.memory', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='persistentVolume', full_name='flyteidl_flink.Resource.persistentVolume', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_RESOURCE_PERSISTENTVOLUME, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=101,
  serialized_end=487,
)


_JOBMANAGER = _descriptor.Descriptor(
  name='JobManager',
  full_name='flyteidl_flink.JobManager',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='resource', full_name='flyteidl_flink.JobManager.resource', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=489,
  serialized_end=545,
)


_TASKMANAGER = _descriptor.Descriptor(
  name='TaskManager',
  full_name='flyteidl_flink.TaskManager',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='resource', full_name='flyteidl_flink.TaskManager.resource', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='replicas', full_name='flyteidl_flink.TaskManager.replicas', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=547,
  serialized_end=622,
)


_FLINKJOB_FLINKPROPERTIESENTRY = _descriptor.Descriptor(
  name='FlinkPropertiesEntry',
  full_name='flyteidl_flink.FlinkJob.FlinkPropertiesEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='flyteidl_flink.FlinkJob.FlinkPropertiesEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='flyteidl_flink.FlinkJob.FlinkPropertiesEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=896,
  serialized_end=950,
)

_FLINKJOB = _descriptor.Descriptor(
  name='FlinkJob',
  full_name='flyteidl_flink.FlinkJob',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='jarFile', full_name='flyteidl_flink.FlinkJob.jarFile', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='mainClass', full_name='flyteidl_flink.FlinkJob.mainClass', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='args', full_name='flyteidl_flink.FlinkJob.args', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='flinkProperties', full_name='flyteidl_flink.FlinkJob.flinkProperties', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='jobManager', full_name='flyteidl_flink.FlinkJob.jobManager', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='taskManager', full_name='flyteidl_flink.FlinkJob.taskManager', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='serviceAccount', full_name='flyteidl_flink.FlinkJob.serviceAccount', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='image', full_name='flyteidl_flink.FlinkJob.image', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_FLINKJOB_FLINKPROPERTIESENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=625,
  serialized_end=950,
)


_JOBEXECUTIONINFO = _descriptor.Descriptor(
  name='JobExecutionInfo',
  full_name='flyteidl_flink.JobExecutionInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl_flink.JobExecutionInfo.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=952,
  serialized_end=982,
)


_JOBMANAGEREXECUTIONINFO = _descriptor.Descriptor(
  name='JobManagerExecutionInfo',
  full_name='flyteidl_flink.JobManagerExecutionInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ingressURLs', full_name='flyteidl_flink.JobManagerExecutionInfo.ingressURLs', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=984,
  serialized_end=1030,
)


_FLINKEXECUTIONINFO = _descriptor.Descriptor(
  name='FlinkExecutionInfo',
  full_name='flyteidl_flink.FlinkExecutionInfo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='job', full_name='flyteidl_flink.FlinkExecutionInfo.job', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='jobManager', full_name='flyteidl_flink.FlinkExecutionInfo.jobManager', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1033,
  serialized_end=1161,
)

_RESOURCE_PERSISTENTVOLUME.fields_by_name['type'].enum_type = _RESOURCE_PERSISTENTVOLUME_TYPE
_RESOURCE_PERSISTENTVOLUME.fields_by_name['size'].message_type = k8s_dot_io_dot_apimachinery_dot_pkg_dot_api_dot_resource_dot_generated__pb2._QUANTITY
_RESOURCE_PERSISTENTVOLUME.containing_type = _RESOURCE
_RESOURCE_PERSISTENTVOLUME_TYPE.containing_type = _RESOURCE_PERSISTENTVOLUME
_RESOURCE.fields_by_name['cpu'].message_type = k8s_dot_io_dot_apimachinery_dot_pkg_dot_api_dot_resource_dot_generated__pb2._QUANTITY
_RESOURCE.fields_by_name['memory'].message_type = k8s_dot_io_dot_apimachinery_dot_pkg_dot_api_dot_resource_dot_generated__pb2._QUANTITY
_RESOURCE.fields_by_name['persistentVolume'].message_type = _RESOURCE_PERSISTENTVOLUME
_JOBMANAGER.fields_by_name['resource'].message_type = _RESOURCE
_TASKMANAGER.fields_by_name['resource'].message_type = _RESOURCE
_FLINKJOB_FLINKPROPERTIESENTRY.containing_type = _FLINKJOB
_FLINKJOB.fields_by_name['flinkProperties'].message_type = _FLINKJOB_FLINKPROPERTIESENTRY
_FLINKJOB.fields_by_name['jobManager'].message_type = _JOBMANAGER
_FLINKJOB.fields_by_name['taskManager'].message_type = _TASKMANAGER
_FLINKEXECUTIONINFO.fields_by_name['job'].message_type = _JOBEXECUTIONINFO
_FLINKEXECUTIONINFO.fields_by_name['jobManager'].message_type = _JOBMANAGEREXECUTIONINFO
DESCRIPTOR.message_types_by_name['Resource'] = _RESOURCE
DESCRIPTOR.message_types_by_name['JobManager'] = _JOBMANAGER
DESCRIPTOR.message_types_by_name['TaskManager'] = _TASKMANAGER
DESCRIPTOR.message_types_by_name['FlinkJob'] = _FLINKJOB
DESCRIPTOR.message_types_by_name['JobExecutionInfo'] = _JOBEXECUTIONINFO
DESCRIPTOR.message_types_by_name['JobManagerExecutionInfo'] = _JOBMANAGEREXECUTIONINFO
DESCRIPTOR.message_types_by_name['FlinkExecutionInfo'] = _FLINKEXECUTIONINFO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Resource = _reflection.GeneratedProtocolMessageType('Resource', (_message.Message,), dict(

  PersistentVolume = _reflection.GeneratedProtocolMessageType('PersistentVolume', (_message.Message,), dict(
    DESCRIPTOR = _RESOURCE_PERSISTENTVOLUME,
    __module__ = 'flyteidl_flink.flink_pb2'
    # @@protoc_insertion_point(class_scope:flyteidl_flink.Resource.PersistentVolume)
    ))
  ,
  DESCRIPTOR = _RESOURCE,
  __module__ = 'flyteidl_flink.flink_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl_flink.Resource)
  ))
_sym_db.RegisterMessage(Resource)
_sym_db.RegisterMessage(Resource.PersistentVolume)

JobManager = _reflection.GeneratedProtocolMessageType('JobManager', (_message.Message,), dict(
  DESCRIPTOR = _JOBMANAGER,
  __module__ = 'flyteidl_flink.flink_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl_flink.JobManager)
  ))
_sym_db.RegisterMessage(JobManager)

TaskManager = _reflection.GeneratedProtocolMessageType('TaskManager', (_message.Message,), dict(
  DESCRIPTOR = _TASKMANAGER,
  __module__ = 'flyteidl_flink.flink_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl_flink.TaskManager)
  ))
_sym_db.RegisterMessage(TaskManager)

FlinkJob = _reflection.GeneratedProtocolMessageType('FlinkJob', (_message.Message,), dict(

  FlinkPropertiesEntry = _reflection.GeneratedProtocolMessageType('FlinkPropertiesEntry', (_message.Message,), dict(
    DESCRIPTOR = _FLINKJOB_FLINKPROPERTIESENTRY,
    __module__ = 'flyteidl_flink.flink_pb2'
    # @@protoc_insertion_point(class_scope:flyteidl_flink.FlinkJob.FlinkPropertiesEntry)
    ))
  ,
  DESCRIPTOR = _FLINKJOB,
  __module__ = 'flyteidl_flink.flink_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl_flink.FlinkJob)
  ))
_sym_db.RegisterMessage(FlinkJob)
_sym_db.RegisterMessage(FlinkJob.FlinkPropertiesEntry)

JobExecutionInfo = _reflection.GeneratedProtocolMessageType('JobExecutionInfo', (_message.Message,), dict(
  DESCRIPTOR = _JOBEXECUTIONINFO,
  __module__ = 'flyteidl_flink.flink_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl_flink.JobExecutionInfo)
  ))
_sym_db.RegisterMessage(JobExecutionInfo)

JobManagerExecutionInfo = _reflection.GeneratedProtocolMessageType('JobManagerExecutionInfo', (_message.Message,), dict(
  DESCRIPTOR = _JOBMANAGEREXECUTIONINFO,
  __module__ = 'flyteidl_flink.flink_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl_flink.JobManagerExecutionInfo)
  ))
_sym_db.RegisterMessage(JobManagerExecutionInfo)

FlinkExecutionInfo = _reflection.GeneratedProtocolMessageType('FlinkExecutionInfo', (_message.Message,), dict(
  DESCRIPTOR = _FLINKEXECUTIONINFO,
  __module__ = 'flyteidl_flink.flink_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl_flink.FlinkExecutionInfo)
  ))
_sym_db.RegisterMessage(FlinkExecutionInfo)


DESCRIPTOR._options = None
_FLINKJOB_FLINKPROPERTIESENTRY._options = None
# @@protoc_insertion_point(module_scope)
