//
// Cosmo Dashboard API
// Manipulate cosmo dashboard resource API

// @generated by protoc-gen-es v0.2.1 with parameter "target=ts"
// @generated from file dashboard/v1alpha1/workspace.proto (package dashboard.v1alpha1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3, protoInt64} from "@bufbuild/protobuf";

/**
 * @generated from message dashboard.v1alpha1.NetworkRule
 */
export class NetworkRule extends Message<NetworkRule> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: int32 port_number = 2;
   */
  portNumber = 0;

  /**
   * @generated from field: string group = 3;
   */
  group = "";

  /**
   * @generated from field: string http_path = 4;
   */
  httpPath = "";

  /**
   * @generated from field: string url = 5;
   */
  url = "";

  /**
   * @generated from field: bool public = 6;
   */
  public = false;

  constructor(data?: PartialMessage<NetworkRule>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.NetworkRule";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "port_number", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "group", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "http_path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "public", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NetworkRule {
    return new NetworkRule().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NetworkRule {
    return new NetworkRule().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NetworkRule {
    return new NetworkRule().fromJsonString(jsonString, options);
  }

  static equals(a: NetworkRule | PlainMessage<NetworkRule> | undefined, b: NetworkRule | PlainMessage<NetworkRule> | undefined): boolean {
    return proto3.util.equals(NetworkRule, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.WorkspaceSpec
 */
export class WorkspaceSpec extends Message<WorkspaceSpec> {
  /**
   * @generated from field: string template = 1;
   */
  template = "";

  /**
   * @generated from field: int64 replicas = 2;
   */
  replicas = protoInt64.zero;

  /**
   * @generated from field: map<string, string> vars = 3;
   */
  vars: { [key: string]: string } = {};

  /**
   * @generated from field: repeated dashboard.v1alpha1.NetworkRule additional_network = 4;
   */
  additionalNetwork: NetworkRule[] = [];

  constructor(data?: PartialMessage<WorkspaceSpec>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.WorkspaceSpec";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "template", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "replicas", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "vars", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 4, name: "additional_network", kind: "message", T: NetworkRule, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkspaceSpec {
    return new WorkspaceSpec().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkspaceSpec {
    return new WorkspaceSpec().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkspaceSpec {
    return new WorkspaceSpec().fromJsonString(jsonString, options);
  }

  static equals(a: WorkspaceSpec | PlainMessage<WorkspaceSpec> | undefined, b: WorkspaceSpec | PlainMessage<WorkspaceSpec> | undefined): boolean {
    return proto3.util.equals(WorkspaceSpec, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.WorkspaceStatus
 */
export class WorkspaceStatus extends Message<WorkspaceStatus> {
  /**
   * @generated from field: string phase = 1;
   */
  phase = "";

  /**
   * @generated from field: string main_url = 2;
   */
  mainUrl = "";

  /**
   * @generated from field: string url_base = 3;
   */
  urlBase = "";

  constructor(data?: PartialMessage<WorkspaceStatus>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.WorkspaceStatus";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "phase", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "main_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "url_base", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkspaceStatus {
    return new WorkspaceStatus().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkspaceStatus {
    return new WorkspaceStatus().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkspaceStatus {
    return new WorkspaceStatus().fromJsonString(jsonString, options);
  }

  static equals(a: WorkspaceStatus | PlainMessage<WorkspaceStatus> | undefined, b: WorkspaceStatus | PlainMessage<WorkspaceStatus> | undefined): boolean {
    return proto3.util.equals(WorkspaceStatus, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.Workspace
 */
export class Workspace extends Message<Workspace> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string owner_id = 2;
   */
  ownerId = "";

  /**
   * @generated from field: dashboard.v1alpha1.WorkspaceSpec spec = 3;
   */
  spec?: WorkspaceSpec;

  /**
   * @generated from field: dashboard.v1alpha1.WorkspaceStatus status = 4;
   */
  status?: WorkspaceStatus;

  constructor(data?: PartialMessage<Workspace>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.Workspace";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "owner_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "spec", kind: "message", T: WorkspaceSpec },
    { no: 4, name: "status", kind: "message", T: WorkspaceStatus },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Workspace {
    return new Workspace().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Workspace {
    return new Workspace().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Workspace {
    return new Workspace().fromJsonString(jsonString, options);
  }

  static equals(a: Workspace | PlainMessage<Workspace> | undefined, b: Workspace | PlainMessage<Workspace> | undefined): boolean {
    return proto3.util.equals(Workspace, a, b);
  }
}
