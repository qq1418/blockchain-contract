# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: contract_execute_logs.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import contract_pb2 as contract__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='contract_execute_logs.proto',
  package='protos',
  syntax='proto3',
  serialized_pb=_b('\n\x1b\x63ontract_execute_logs.proto\x12\x06protos\x1a\x0e\x63ontract.proto\"\xb2\x02\n\x12\x43ontractExecuteLog\x12\x16\n\x0e\x43ontractHashId\x18\x01 \x01(\t\x12\x0e\n\x06TaskId\x18\x02 \x01(\t\x12\x11\n\tTimestamp\x18\x03 \x01(\t\x12\x0f\n\x07\x43\x61ption\x18\x04 \x01(\t\x12\r\n\x05\x43name\x18\x05 \x01(\t\x12\r\n\x05\x43type\x18\x06 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x07 \x01(\t\x12\x44\n\rMetaAttribute\x18\x08 \x03(\x0b\x32-.protos.ContractExecuteLog.MetaAttributeEntry\x12\r\n\x05State\x18\t \x01(\t\x12\x12\n\nCreateTime\x18\n \x01(\t\x1a\x34\n\x12MetaAttributeEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"J\n\x16\x43ontractExecuteLogList\x12\x30\n\x0c\x43ontractLogs\x18\x01 \x03(\x0b\x32\x1a.protos.ContractExecuteLogB+\n\x14\x63om.uniledger.protosB\x13\x43ontractExecuteLogsb\x06proto3')
  ,
  dependencies=[contract__pb2.DESCRIPTOR,])




_CONTRACTEXECUTELOG_METAATTRIBUTEENTRY = _descriptor.Descriptor(
  name='MetaAttributeEntry',
  full_name='protos.ContractExecuteLog.MetaAttributeEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='protos.ContractExecuteLog.MetaAttributeEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='value', full_name='protos.ContractExecuteLog.MetaAttributeEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=_descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001')),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=310,
  serialized_end=362,
)

_CONTRACTEXECUTELOG = _descriptor.Descriptor(
  name='ContractExecuteLog',
  full_name='protos.ContractExecuteLog',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ContractHashId', full_name='protos.ContractExecuteLog.ContractHashId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='TaskId', full_name='protos.ContractExecuteLog.TaskId', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Timestamp', full_name='protos.ContractExecuteLog.Timestamp', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Caption', full_name='protos.ContractExecuteLog.Caption', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Cname', full_name='protos.ContractExecuteLog.Cname', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Ctype', full_name='protos.ContractExecuteLog.Ctype', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='protos.ContractExecuteLog.Description', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='MetaAttribute', full_name='protos.ContractExecuteLog.MetaAttribute', index=7,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='State', full_name='protos.ContractExecuteLog.State', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='CreateTime', full_name='protos.ContractExecuteLog.CreateTime', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[_CONTRACTEXECUTELOG_METAATTRIBUTEENTRY, ],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=56,
  serialized_end=362,
)


_CONTRACTEXECUTELOGLIST = _descriptor.Descriptor(
  name='ContractExecuteLogList',
  full_name='protos.ContractExecuteLogList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ContractLogs', full_name='protos.ContractExecuteLogList.ContractLogs', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=364,
  serialized_end=438,
)

_CONTRACTEXECUTELOG_METAATTRIBUTEENTRY.containing_type = _CONTRACTEXECUTELOG
_CONTRACTEXECUTELOG.fields_by_name['MetaAttribute'].message_type = _CONTRACTEXECUTELOG_METAATTRIBUTEENTRY
_CONTRACTEXECUTELOGLIST.fields_by_name['ContractLogs'].message_type = _CONTRACTEXECUTELOG
DESCRIPTOR.message_types_by_name['ContractExecuteLog'] = _CONTRACTEXECUTELOG
DESCRIPTOR.message_types_by_name['ContractExecuteLogList'] = _CONTRACTEXECUTELOGLIST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContractExecuteLog = _reflection.GeneratedProtocolMessageType('ContractExecuteLog', (_message.Message,), dict(

  MetaAttributeEntry = _reflection.GeneratedProtocolMessageType('MetaAttributeEntry', (_message.Message,), dict(
    DESCRIPTOR = _CONTRACTEXECUTELOG_METAATTRIBUTEENTRY,
    __module__ = 'contract_execute_logs_pb2'
    # @@protoc_insertion_point(class_scope:protos.ContractExecuteLog.MetaAttributeEntry)
    ))
  ,
  DESCRIPTOR = _CONTRACTEXECUTELOG,
  __module__ = 'contract_execute_logs_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractExecuteLog)
  ))
_sym_db.RegisterMessage(ContractExecuteLog)
_sym_db.RegisterMessage(ContractExecuteLog.MetaAttributeEntry)

ContractExecuteLogList = _reflection.GeneratedProtocolMessageType('ContractExecuteLogList', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACTEXECUTELOGLIST,
  __module__ = 'contract_execute_logs_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractExecuteLogList)
  ))
_sym_db.RegisterMessage(ContractExecuteLogList)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\024com.uniledger.protosB\023ContractExecuteLogs'))
_CONTRACTEXECUTELOG_METAATTRIBUTEENTRY.has_options = True
_CONTRACTEXECUTELOG_METAATTRIBUTEENTRY._options = _descriptor._ParseOptions(descriptor_pb2.MessageOptions(), _b('8\001'))
# @@protoc_insertion_point(module_scope)
