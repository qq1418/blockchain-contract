# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: contract_list.proto

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
  name='contract_list.proto',
  package='protos',
  syntax='proto3',
  serialized_pb=_b('\n\x13\x63ontract_list.proto\x12\x06protos\x1a\x0e\x63ontract.proto\"3\n\x0c\x43ontractList\x12#\n\tcontracts\x18\x01 \x03(\x0b\x32\x10.protos.ContractB)\n\x14\x63om.uniledger.protosB\x11ProtoContractListb\x06proto3')
  ,
  dependencies=[contract__pb2.DESCRIPTOR,])




_CONTRACTLIST = _descriptor.Descriptor(
  name='ContractList',
  full_name='protos.ContractList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='contracts', full_name='protos.ContractList.contracts', index=0,
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
  serialized_start=47,
  serialized_end=98,
)

_CONTRACTLIST.fields_by_name['contracts'].message_type = contract__pb2._CONTRACT
DESCRIPTOR.message_types_by_name['ContractList'] = _CONTRACTLIST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContractList = _reflection.GeneratedProtocolMessageType('ContractList', (_message.Message,), dict(
  DESCRIPTOR = _CONTRACTLIST,
  __module__ = 'contract_list_pb2'
  # @@protoc_insertion_point(class_scope:protos.ContractList)
  ))
_sym_db.RegisterMessage(ContractList)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('\n\024com.uniledger.protosB\021ProtoContractList'))
# @@protoc_insertion_point(module_scope)
