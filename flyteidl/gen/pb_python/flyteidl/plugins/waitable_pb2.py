# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/plugins/waitable.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import execution_pb2 as flyteidl_dot_core_dot_execution__pb2
from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1f\x66lyteidl/plugins/waitable.proto\x12\x10\x66lyteidl.plugins\x1a\x1d\x66lyteidl/core/execution.proto\x1a\x1e\x66lyteidl/core/identifier.proto\"\xb3\x01\n\x08Waitable\x12H\n\nwf_exec_id\x18\x01 \x01(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifierR\x08wfExecId\x12<\n\x05phase\x18\x02 \x01(\x0e\x32&.flyteidl.core.WorkflowExecution.PhaseR\x05phase\x12\x1f\n\x0bworkflow_id\x18\x03 \x01(\tR\nworkflowIdB\xbf\x01\n\x14\x63om.flyteidl.pluginsB\rWaitableProtoP\x01Z7github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/plugins\xa2\x02\x03\x46PX\xaa\x02\x10\x46lyteidl.Plugins\xca\x02\x10\x46lyteidl\\Plugins\xe2\x02\x1c\x46lyteidl\\Plugins\\GPBMetadata\xea\x02\x11\x46lyteidl::Pluginsb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.plugins.waitable_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\024com.flyteidl.pluginsB\rWaitableProtoP\001Z7github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/plugins\242\002\003FPX\252\002\020Flyteidl.Plugins\312\002\020Flyteidl\\Plugins\342\002\034Flyteidl\\Plugins\\GPBMetadata\352\002\021Flyteidl::Plugins'
  _globals['_WAITABLE']._serialized_start=117
  _globals['_WAITABLE']._serialized_end=296
# @@protoc_insertion_point(module_scope)
