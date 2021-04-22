.. _api_file_validate/validate.proto:

validate.proto
=======================

.. _api_msg_validate.FieldRules:

validate.FieldRules
-------------------

`[validate.FieldRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L33>`_

FieldRules encapsulates the rules for each type of field. Depending on the
field, the correct set should be used to ensure proper validations.

.. code-block:: json

  {
    "message": "{...}",
    "float": "{...}",
    "double": "{...}",
    "int32": "{...}",
    "int64": "{...}",
    "uint32": "{...}",
    "uint64": "{...}",
    "sint32": "{...}",
    "sint64": "{...}",
    "fixed32": "{...}",
    "fixed64": "{...}",
    "sfixed32": "{...}",
    "sfixed64": "{...}",
    "bool": "{...}",
    "string": "{...}",
    "bytes": "{...}",
    "enum": "{...}",
    "repeated": "{...}",
    "map": "{...}",
    "any": "{...}",
    "duration": "{...}",
    "timestamp": "{...}"
  }

.. _api_field_validate.FieldRules.message:

message
  (:ref:`validate.MessageRules <api_msg_validate.MessageRules>`) 
  
.. _api_field_validate.FieldRules.float:

float
  (:ref:`validate.FloatRules <api_msg_validate.FloatRules>`) Scalar Field Types
  
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.double:

double
  (:ref:`validate.DoubleRules <api_msg_validate.DoubleRules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.int32:

int32
  (:ref:`validate.Int32Rules <api_msg_validate.Int32Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.int64:

int64
  (:ref:`validate.Int64Rules <api_msg_validate.Int64Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.uint32:

uint32
  (:ref:`validate.UInt32Rules <api_msg_validate.UInt32Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.uint64:

uint64
  (:ref:`validate.UInt64Rules <api_msg_validate.UInt64Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.sint32:

sint32
  (:ref:`validate.SInt32Rules <api_msg_validate.SInt32Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.sint64:

sint64
  (:ref:`validate.SInt64Rules <api_msg_validate.SInt64Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.fixed32:

fixed32
  (:ref:`validate.Fixed32Rules <api_msg_validate.Fixed32Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.fixed64:

fixed64
  (:ref:`validate.Fixed64Rules <api_msg_validate.Fixed64Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.sfixed32:

sfixed32
  (:ref:`validate.SFixed32Rules <api_msg_validate.SFixed32Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.sfixed64:

sfixed64
  (:ref:`validate.SFixed64Rules <api_msg_validate.SFixed64Rules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.bool:

bool
  (:ref:`validate.BoolRules <api_msg_validate.BoolRules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.string:

string
  (:ref:`validate.StringRules <api_msg_validate.StringRules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.bytes:

bytes
  (:ref:`validate.BytesRules <api_msg_validate.BytesRules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.enum:

enum
  (:ref:`validate.EnumRules <api_msg_validate.EnumRules>`) Complex Field Types
  
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.repeated:

repeated
  (:ref:`validate.RepeatedRules <api_msg_validate.RepeatedRules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.map:

map
  (:ref:`validate.MapRules <api_msg_validate.MapRules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.any:

any
  (:ref:`validate.AnyRules <api_msg_validate.AnyRules>`) Well-Known Field Types
  
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.duration:

duration
  (:ref:`validate.DurationRules <api_msg_validate.DurationRules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  
.. _api_field_validate.FieldRules.timestamp:

timestamp
  (:ref:`validate.TimestampRules <api_msg_validate.TimestampRules>`) 
  
  
  Only one of :ref:`float <api_field_validate.FieldRules.float>`, :ref:`double <api_field_validate.FieldRules.double>`, :ref:`int32 <api_field_validate.FieldRules.int32>`, :ref:`int64 <api_field_validate.FieldRules.int64>`, :ref:`uint32 <api_field_validate.FieldRules.uint32>`, :ref:`uint64 <api_field_validate.FieldRules.uint64>`, :ref:`sint32 <api_field_validate.FieldRules.sint32>`, :ref:`sint64 <api_field_validate.FieldRules.sint64>`, :ref:`fixed32 <api_field_validate.FieldRules.fixed32>`, :ref:`fixed64 <api_field_validate.FieldRules.fixed64>`, :ref:`sfixed32 <api_field_validate.FieldRules.sfixed32>`, :ref:`sfixed64 <api_field_validate.FieldRules.sfixed64>`, :ref:`bool <api_field_validate.FieldRules.bool>`, :ref:`string <api_field_validate.FieldRules.string>`, :ref:`bytes <api_field_validate.FieldRules.bytes>`, :ref:`enum <api_field_validate.FieldRules.enum>`, :ref:`repeated <api_field_validate.FieldRules.repeated>`, :ref:`map <api_field_validate.FieldRules.map>`, :ref:`any <api_field_validate.FieldRules.any>`, :ref:`duration <api_field_validate.FieldRules.duration>`, :ref:`timestamp <api_field_validate.FieldRules.timestamp>` may be set.
  


.. _api_msg_validate.FloatRules:

validate.FloatRules
-------------------

`[validate.FloatRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L66>`_

FloatRules describes the constraints applied to `float` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.FloatRules.const:

const
  (`float <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.FloatRules.lt:

lt
  (`float <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.FloatRules.lte:

lte
  (`float <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.FloatRules.gt:

gt
  (`float <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.FloatRules.gte:

gte
  (`float <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.FloatRules.in:

in
  (`float <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.FloatRules.not_in:

not_in
  (`float <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.DoubleRules:

validate.DoubleRules
--------------------

`[validate.DoubleRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L98>`_

DoubleRules describes the constraints applied to `double` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.DoubleRules.const:

const
  (`double <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.DoubleRules.lt:

lt
  (`double <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.DoubleRules.lte:

lte
  (`double <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.DoubleRules.gt:

gt
  (`double <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.DoubleRules.gte:

gte
  (`double <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.DoubleRules.in:

in
  (`double <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.DoubleRules.not_in:

not_in
  (`double <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.Int32Rules:

validate.Int32Rules
-------------------

`[validate.Int32Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L130>`_

Int32Rules describes the constraints applied to `int32` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.Int32Rules.const:

const
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.Int32Rules.lt:

lt
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.Int32Rules.lte:

lte
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.Int32Rules.gt:

gt
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.Int32Rules.gte:

gte
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.Int32Rules.in:

in
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.Int32Rules.not_in:

not_in
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.Int64Rules:

validate.Int64Rules
-------------------

`[validate.Int64Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L162>`_

Int64Rules describes the constraints applied to `int64` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.Int64Rules.const:

const
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.Int64Rules.lt:

lt
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.Int64Rules.lte:

lte
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.Int64Rules.gt:

gt
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.Int64Rules.gte:

gte
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.Int64Rules.in:

in
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.Int64Rules.not_in:

not_in
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.UInt32Rules:

validate.UInt32Rules
--------------------

`[validate.UInt32Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L194>`_

UInt32Rules describes the constraints applied to `uint32` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.UInt32Rules.const:

const
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.UInt32Rules.lt:

lt
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.UInt32Rules.lte:

lte
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.UInt32Rules.gt:

gt
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.UInt32Rules.gte:

gte
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.UInt32Rules.in:

in
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.UInt32Rules.not_in:

not_in
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.UInt64Rules:

validate.UInt64Rules
--------------------

`[validate.UInt64Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L226>`_

UInt64Rules describes the constraints applied to `uint64` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.UInt64Rules.const:

const
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.UInt64Rules.lt:

lt
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.UInt64Rules.lte:

lte
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.UInt64Rules.gt:

gt
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.UInt64Rules.gte:

gte
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.UInt64Rules.in:

in
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.UInt64Rules.not_in:

not_in
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.SInt32Rules:

validate.SInt32Rules
--------------------

`[validate.SInt32Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L258>`_

SInt32Rules describes the constraints applied to `sint32` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.SInt32Rules.const:

const
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.SInt32Rules.lt:

lt
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.SInt32Rules.lte:

lte
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.SInt32Rules.gt:

gt
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.SInt32Rules.gte:

gte
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.SInt32Rules.in:

in
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.SInt32Rules.not_in:

not_in
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.SInt64Rules:

validate.SInt64Rules
--------------------

`[validate.SInt64Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L290>`_

SInt64Rules describes the constraints applied to `sint64` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.SInt64Rules.const:

const
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.SInt64Rules.lt:

lt
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.SInt64Rules.lte:

lte
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.SInt64Rules.gt:

gt
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.SInt64Rules.gte:

gte
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.SInt64Rules.in:

in
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.SInt64Rules.not_in:

not_in
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.Fixed32Rules:

validate.Fixed32Rules
---------------------

`[validate.Fixed32Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L322>`_

Fixed32Rules describes the constraints applied to `fixed32` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.Fixed32Rules.const:

const
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.Fixed32Rules.lt:

lt
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.Fixed32Rules.lte:

lte
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.Fixed32Rules.gt:

gt
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.Fixed32Rules.gte:

gte
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.Fixed32Rules.in:

in
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.Fixed32Rules.not_in:

not_in
  (`uint32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.Fixed64Rules:

validate.Fixed64Rules
---------------------

`[validate.Fixed64Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L354>`_

Fixed64Rules describes the constraints applied to `fixed64` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.Fixed64Rules.const:

const
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.Fixed64Rules.lt:

lt
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.Fixed64Rules.lte:

lte
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.Fixed64Rules.gt:

gt
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.Fixed64Rules.gte:

gte
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.Fixed64Rules.in:

in
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.Fixed64Rules.not_in:

not_in
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.SFixed32Rules:

validate.SFixed32Rules
----------------------

`[validate.SFixed32Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L386>`_

SFixed32Rules describes the constraints applied to `sfixed32` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.SFixed32Rules.const:

const
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.SFixed32Rules.lt:

lt
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.SFixed32Rules.lte:

lte
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.SFixed32Rules.gt:

gt
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.SFixed32Rules.gte:

gte
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.SFixed32Rules.in:

in
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.SFixed32Rules.not_in:

not_in
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.SFixed64Rules:

validate.SFixed64Rules
----------------------

`[validate.SFixed64Rules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L418>`_

SFixed64Rules describes the constraints applied to `sfixed64` values

.. code-block:: json

  {
    "const": "...",
    "lt": "...",
    "lte": "...",
    "gt": "...",
    "gte": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.SFixed64Rules.const:

const
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.SFixed64Rules.lt:

lt
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.SFixed64Rules.lte:

lte
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Lte specifies that this field must be less than or equal to the
  specified value, inclusive
  
  
.. _api_field_validate.SFixed64Rules.gt:

gt
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gt specifies that this field must be greater than the specified value,
  exclusive. If the value of Gt is larger than a specified Lt or Lte, the
  range is reversed.
  
  
.. _api_field_validate.SFixed64Rules.gte:

gte
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Gte specifies that this field must be greater than or equal to the
  specified value, inclusive. If the value of Gte is larger than a
  specified Lt or Lte, the range is reversed.
  
  
.. _api_field_validate.SFixed64Rules.in:

in
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.SFixed64Rules.not_in:

not_in
  (`int64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.BoolRules:

validate.BoolRules
------------------

`[validate.BoolRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L450>`_

BoolRules describes the constraints applied to `bool` values

.. code-block:: json

  {
    "const": "..."
  }

.. _api_field_validate.BoolRules.const:

const
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  


.. _api_msg_validate.StringRules:

validate.StringRules
--------------------

`[validate.StringRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L456>`_

StringRules describe the constraints applied to `string` values

.. code-block:: json

  {
    "const": "...",
    "len": "...",
    "min_len": "...",
    "max_len": "...",
    "len_bytes": "...",
    "min_bytes": "...",
    "max_bytes": "...",
    "pattern": "...",
    "prefix": "...",
    "suffix": "...",
    "contains": "...",
    "not_contains": "...",
    "in": [],
    "not_in": [],
    "email": "...",
    "hostname": "...",
    "ip": "...",
    "ipv4": "...",
    "ipv6": "...",
    "uri": "...",
    "uri_ref": "...",
    "address": "...",
    "uuid": "...",
    "well_known_regex": "...",
    "strict": "..."
  }

.. _api_field_validate.StringRules.const:

const
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.StringRules.len:

len
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Len specifies that this field must be the specified number of
  characters (Unicode code points). Note that the number of
  characters may differ from the number of bytes in the string.
  
  
.. _api_field_validate.StringRules.min_len:

min_len
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MinLen specifies that this field must be the specified number of
  characters (Unicode code points) at a minimum. Note that the number of
  characters may differ from the number of bytes in the string.
  
  
.. _api_field_validate.StringRules.max_len:

max_len
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MaxLen specifies that this field must be the specified number of
  characters (Unicode code points) at a maximum. Note that the number of
  characters may differ from the number of bytes in the string.
  
  
.. _api_field_validate.StringRules.len_bytes:

len_bytes
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) LenBytes specifies that this field must be the specified number of bytes
  at a minimum
  
  
.. _api_field_validate.StringRules.min_bytes:

min_bytes
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MinBytes specifies that this field must be the specified number of bytes
  at a minimum
  
  
.. _api_field_validate.StringRules.max_bytes:

max_bytes
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MaxBytes specifies that this field must be the specified number of bytes
  at a maximum
  
  
.. _api_field_validate.StringRules.pattern:

pattern
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Pattern specifes that this field must match against the specified
  regular expression (RE2 syntax). The included expression should elide
  any delimiters.
  
  
.. _api_field_validate.StringRules.prefix:

prefix
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Prefix specifies that this field must have the specified substring at
  the beginning of the string.
  
  
.. _api_field_validate.StringRules.suffix:

suffix
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Suffix specifies that this field must have the specified substring at
  the end of the string.
  
  
.. _api_field_validate.StringRules.contains:

contains
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Contains specifies that this field must have the specified substring
  anywhere in the string.
  
  
.. _api_field_validate.StringRules.not_contains:

not_contains
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotContains specifies that this field cannot have the specified substring
  anywhere in the string.
  
  
.. _api_field_validate.StringRules.in:

in
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.StringRules.not_in:

not_in
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  
.. _api_field_validate.StringRules.email:

email
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Email specifies that the field must be a valid email address as
  defined by RFC 5322
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.hostname:

hostname
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Hostname specifies that the field must be a valid hostname as
  defined by RFC 1034. This constraint does not support
  internationalized domain names (IDNs).
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.ip:

ip
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Ip specifies that the field must be a valid IP (v4 or v6) address.
  Valid IPv6 addresses should not include surrounding square brackets.
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.ipv4:

ipv4
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Ipv4 specifies that the field must be a valid IPv4 address.
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.ipv6:

ipv6
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Ipv6 specifies that the field must be a valid IPv6 address. Valid
  IPv6 addresses should not include surrounding square brackets.
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.uri:

uri
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Uri specifies that the field must be a valid, absolute URI as defined
  by RFC 3986
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.uri_ref:

uri_ref
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) UriRef specifies that the field must be a valid URI as defined by RFC
  3986 and may be relative or absolute.
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.address:

address
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Address specifies that the field must be either a valid hostname as
  defined by RFC 1034 (which does not support internationalized domain
  names or IDNs), or it can be a valid IP (v4 or v6).
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.uuid:

uuid
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Uuid specifies that the field must be a valid UUID as defined by
  RFC 4122
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.well_known_regex:

well_known_regex
  (:ref:`validate.KnownRegex <api_enum_validate.KnownRegex>`) WellKnownRegex specifies a common well known pattern defined as a regex.
  
  WellKnown rules provide advanced constraints against common string
  patterns
  
  
  Only one of :ref:`email <api_field_validate.StringRules.email>`, :ref:`hostname <api_field_validate.StringRules.hostname>`, :ref:`ip <api_field_validate.StringRules.ip>`, :ref:`ipv4 <api_field_validate.StringRules.ipv4>`, :ref:`ipv6 <api_field_validate.StringRules.ipv6>`, :ref:`uri <api_field_validate.StringRules.uri>`, :ref:`uri_ref <api_field_validate.StringRules.uri_ref>`, :ref:`address <api_field_validate.StringRules.address>`, :ref:`uuid <api_field_validate.StringRules.uuid>`, :ref:`well_known_regex <api_field_validate.StringRules.well_known_regex>` may be set.
  
.. _api_field_validate.StringRules.strict:

strict
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) This applies to regexes HTTP_HEADER_NAME and HTTP_HEADER_VALUE to enable
  strict header validation.
  By default, this is true, and HTTP header validations are RFC-compliant.
  Setting to false will enable a looser validations that only disallows
  \r\n\0 characters, which can be used to bypass header matching rules.
  
  


.. _api_msg_validate.BytesRules:

validate.BytesRules
-------------------

`[validate.BytesRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L580>`_

BytesRules describe the constraints applied to `bytes` values

.. code-block:: json

  {
    "const": "...",
    "len": "...",
    "min_len": "...",
    "max_len": "...",
    "pattern": "...",
    "prefix": "...",
    "suffix": "...",
    "contains": "...",
    "in": [],
    "not_in": [],
    "ip": "...",
    "ipv4": "...",
    "ipv6": "..."
  }

.. _api_field_validate.BytesRules.const:

const
  (`bytes <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.BytesRules.len:

len
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Len specifies that this field must be the specified number of bytes
  
  
.. _api_field_validate.BytesRules.min_len:

min_len
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MinLen specifies that this field must be the specified number of bytes
  at a minimum
  
  
.. _api_field_validate.BytesRules.max_len:

max_len
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MaxLen specifies that this field must be the specified number of bytes
  at a maximum
  
  
.. _api_field_validate.BytesRules.pattern:

pattern
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Pattern specifes that this field must match against the specified
  regular expression (RE2 syntax). The included expression should elide
  any delimiters.
  
  
.. _api_field_validate.BytesRules.prefix:

prefix
  (`bytes <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Prefix specifies that this field must have the specified bytes at the
  beginning of the string.
  
  
.. _api_field_validate.BytesRules.suffix:

suffix
  (`bytes <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Suffix specifies that this field must have the specified bytes at the
  end of the string.
  
  
.. _api_field_validate.BytesRules.contains:

contains
  (`bytes <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Contains specifies that this field must have the specified bytes
  anywhere in the string.
  
  
.. _api_field_validate.BytesRules.in:

in
  (`bytes <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.BytesRules.not_in:

not_in
  (`bytes <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  
.. _api_field_validate.BytesRules.ip:

ip
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Ip specifies that the field must be a valid IP (v4 or v6) address in
  byte format
  
  WellKnown rules provide advanced constraints against common byte
  patterns
  
  
  Only one of :ref:`ip <api_field_validate.BytesRules.ip>`, :ref:`ipv4 <api_field_validate.BytesRules.ipv4>`, :ref:`ipv6 <api_field_validate.BytesRules.ipv6>` may be set.
  
.. _api_field_validate.BytesRules.ipv4:

ipv4
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Ipv4 specifies that the field must be a valid IPv4 address in byte
  format
  
  WellKnown rules provide advanced constraints against common byte
  patterns
  
  
  Only one of :ref:`ip <api_field_validate.BytesRules.ip>`, :ref:`ipv4 <api_field_validate.BytesRules.ipv4>`, :ref:`ipv6 <api_field_validate.BytesRules.ipv6>` may be set.
  
.. _api_field_validate.BytesRules.ipv6:

ipv6
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Ipv6 specifies that the field must be a valid IPv6 address in byte
  format
  
  WellKnown rules provide advanced constraints against common byte
  patterns
  
  
  Only one of :ref:`ip <api_field_validate.BytesRules.ip>`, :ref:`ipv4 <api_field_validate.BytesRules.ipv4>`, :ref:`ipv6 <api_field_validate.BytesRules.ipv6>` may be set.
  


.. _api_msg_validate.EnumRules:

validate.EnumRules
------------------

`[validate.EnumRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L638>`_

EnumRules describe the constraints applied to enum values

.. code-block:: json

  {
    "const": "...",
    "defined_only": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.EnumRules.const:

const
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.EnumRules.defined_only:

defined_only
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) DefinedOnly specifies that this field must be only one of the defined
  values for this enum, failing on any undefined value.
  
  
.. _api_field_validate.EnumRules.in:

in
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.EnumRules.not_in:

not_in
  (`int32 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.MessageRules:

validate.MessageRules
---------------------

`[validate.MessageRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L657>`_

MessageRules describe the constraints applied to embedded message values.
For message-type fields, validation is performed recursively.

.. code-block:: json

  {
    "skip": "...",
    "required": "..."
  }

.. _api_field_validate.MessageRules.skip:

skip
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Skip specifies that the validation rules of this field should not be
  evaluated
  
  
.. _api_field_validate.MessageRules.required:

required
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Required specifies that this field must be set
  
  


.. _api_msg_validate.RepeatedRules:

validate.RepeatedRules
----------------------

`[validate.RepeatedRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L667>`_

RepeatedRules describe the constraints applied to `repeated` values

.. code-block:: json

  {
    "min_items": "...",
    "max_items": "...",
    "unique": "...",
    "items": "{...}"
  }

.. _api_field_validate.RepeatedRules.min_items:

min_items
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MinItems specifies that this field must have the specified number of
  items at a minimum
  
  
.. _api_field_validate.RepeatedRules.max_items:

max_items
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MaxItems specifies that this field must have the specified number of
  items at a maximum
  
  
.. _api_field_validate.RepeatedRules.unique:

unique
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Unique specifies that all elements in this field must be unique. This
  contraint is only applicable to scalar and enum types (messages are not
  supported).
  
  
.. _api_field_validate.RepeatedRules.items:

items
  (:ref:`validate.FieldRules <api_msg_validate.FieldRules>`) Items specifies the contraints to be applied to each item in the field.
  Repeated message fields will still execute validation against each item
  unless skip is specified here.
  
  


.. _api_msg_validate.MapRules:

validate.MapRules
-----------------

`[validate.MapRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L688>`_

MapRules describe the constraints applied to `map` values

.. code-block:: json

  {
    "min_pairs": "...",
    "max_pairs": "...",
    "no_sparse": "...",
    "keys": "{...}",
    "values": "{...}"
  }

.. _api_field_validate.MapRules.min_pairs:

min_pairs
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MinPairs specifies that this field must have the specified number of
  KVs at a minimum
  
  
.. _api_field_validate.MapRules.max_pairs:

max_pairs
  (`uint64 <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) MaxPairs specifies that this field must have the specified number of
  KVs at a maximum
  
  
.. _api_field_validate.MapRules.no_sparse:

no_sparse
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NoSparse specifies values in this field cannot be unset. This only
  applies to map's with message value types.
  
  
.. _api_field_validate.MapRules.keys:

keys
  (:ref:`validate.FieldRules <api_msg_validate.FieldRules>`) Keys specifies the constraints to be applied to each key in the field.
  
  
.. _api_field_validate.MapRules.values:

values
  (:ref:`validate.FieldRules <api_msg_validate.FieldRules>`) Values specifies the constraints to be applied to the value of each key
  in the field. Message values will still have their validations evaluated
  unless skip is specified here.
  
  


.. _api_msg_validate.AnyRules:

validate.AnyRules
-----------------

`[validate.AnyRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L712>`_

AnyRules describe constraints applied exclusively to the
`google.protobuf.Any` well-known type

.. code-block:: json

  {
    "required": "...",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.AnyRules.required:

required
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Required specifies that this field must be set
  
  
.. _api_field_validate.AnyRules.in:

in
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) In specifies that this field's `type_url` must be equal to one of the
  specified values.
  
  
.. _api_field_validate.AnyRules.not_in:

not_in
  (`string <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) NotIn specifies that this field's `type_url` must not be equal to any of
  the specified values.
  
  


.. _api_msg_validate.DurationRules:

validate.DurationRules
----------------------

`[validate.DurationRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L727>`_

DurationRules describe the constraints applied exclusively to the
`google.protobuf.Duration` well-known type

.. code-block:: json

  {
    "required": "...",
    "const": "{...}",
    "lt": "{...}",
    "lte": "{...}",
    "gt": "{...}",
    "gte": "{...}",
    "in": [],
    "not_in": []
  }

.. _api_field_validate.DurationRules.required:

required
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Required specifies that this field must be set
  
  
.. _api_field_validate.DurationRules.const:

const
  (:ref:`google.protobuf.Duration <api_msg_google.protobuf.Duration>`) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.DurationRules.lt:

lt
  (:ref:`google.protobuf.Duration <api_msg_google.protobuf.Duration>`) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.DurationRules.lte:

lte
  (:ref:`google.protobuf.Duration <api_msg_google.protobuf.Duration>`) Lt specifies that this field must be less than the specified value,
  inclusive
  
  
.. _api_field_validate.DurationRules.gt:

gt
  (:ref:`google.protobuf.Duration <api_msg_google.protobuf.Duration>`) Gt specifies that this field must be greater than the specified value,
  exclusive
  
  
.. _api_field_validate.DurationRules.gte:

gte
  (:ref:`google.protobuf.Duration <api_msg_google.protobuf.Duration>`) Gte specifies that this field must be greater than the specified value,
  inclusive
  
  
.. _api_field_validate.DurationRules.in:

in
  (:ref:`google.protobuf.Duration <api_msg_google.protobuf.Duration>`) In specifies that this field must be equal to one of the specified
  values
  
  
.. _api_field_validate.DurationRules.not_in:

not_in
  (:ref:`google.protobuf.Duration <api_msg_google.protobuf.Duration>`) NotIn specifies that this field cannot be equal to one of the specified
  values
  
  


.. _api_msg_validate.TimestampRules:

validate.TimestampRules
-----------------------

`[validate.TimestampRules proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L761>`_

TimestampRules describe the constraints applied exclusively to the
`google.protobuf.Timestamp` well-known type

.. code-block:: json

  {
    "required": "...",
    "const": "{...}",
    "lt": "{...}",
    "lte": "{...}",
    "gt": "{...}",
    "gte": "{...}",
    "lt_now": "...",
    "gt_now": "...",
    "within": "{...}"
  }

.. _api_field_validate.TimestampRules.required:

required
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) Required specifies that this field must be set
  
  
.. _api_field_validate.TimestampRules.const:

const
  (:ref:`google.protobuf.Timestamp <api_msg_google.protobuf.Timestamp>`) Const specifies that this field must be exactly the specified value
  
  
.. _api_field_validate.TimestampRules.lt:

lt
  (:ref:`google.protobuf.Timestamp <api_msg_google.protobuf.Timestamp>`) Lt specifies that this field must be less than the specified value,
  exclusive
  
  
.. _api_field_validate.TimestampRules.lte:

lte
  (:ref:`google.protobuf.Timestamp <api_msg_google.protobuf.Timestamp>`) Lte specifies that this field must be less than the specified value,
  inclusive
  
  
.. _api_field_validate.TimestampRules.gt:

gt
  (:ref:`google.protobuf.Timestamp <api_msg_google.protobuf.Timestamp>`) Gt specifies that this field must be greater than the specified value,
  exclusive
  
  
.. _api_field_validate.TimestampRules.gte:

gte
  (:ref:`google.protobuf.Timestamp <api_msg_google.protobuf.Timestamp>`) Gte specifies that this field must be greater than the specified value,
  inclusive
  
  
.. _api_field_validate.TimestampRules.lt_now:

lt_now
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) LtNow specifies that this must be less than the current time. LtNow
  can only be used with the Within rule.
  
  
.. _api_field_validate.TimestampRules.gt_now:

gt_now
  (`bool <https://developers.google.com/protocol-buffers/docs/proto#scalar>`_) GtNow specifies that this must be greater than the current time. GtNow
  can only be used with the Within rule.
  
  
.. _api_field_validate.TimestampRules.within:

within
  (:ref:`google.protobuf.Duration <api_msg_google.protobuf.Duration>`) Within specifies that this field must be within this duration of the
  current time. This constraint can be used alone or with the LtNow and
  GtNow rules.
  
  

.. _api_enum_validate.KnownRegex:

Enum validate.KnownRegex
------------------------

`[validate.KnownRegex proto] <https://github.com/lyft/flyteidl/blob/master/protos/validate/validate.proto#L569>`_

WellKnownRegex contain some well-known patterns.

.. _api_enum_value_validate.KnownRegex.UNKNOWN:

UNKNOWN
  *(DEFAULT)* ⁣
  
.. _api_enum_value_validate.KnownRegex.HTTP_HEADER_NAME:

HTTP_HEADER_NAME
  ⁣HTTP header name as defined by RFC 7230.
  
  
.. _api_enum_value_validate.KnownRegex.HTTP_HEADER_VALUE:

HTTP_HEADER_VALUE
  ⁣HTTP header value as defined by RFC 7230.
  
  
