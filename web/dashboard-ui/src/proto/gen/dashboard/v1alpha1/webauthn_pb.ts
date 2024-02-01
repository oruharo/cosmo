//
//WebAuthn protobuf

// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file dashboard/v1alpha1/webauthn.proto (package dashboard.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message dashboard.v1alpha1.BeginRegistrationRequest
 */
export class BeginRegistrationRequest extends Message<BeginRegistrationRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  constructor(data?: PartialMessage<BeginRegistrationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.BeginRegistrationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BeginRegistrationRequest {
    return new BeginRegistrationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BeginRegistrationRequest {
    return new BeginRegistrationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BeginRegistrationRequest {
    return new BeginRegistrationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: BeginRegistrationRequest | PlainMessage<BeginRegistrationRequest> | undefined, b: BeginRegistrationRequest | PlainMessage<BeginRegistrationRequest> | undefined): boolean {
    return proto3.util.equals(BeginRegistrationRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.BeginRegistrationResponse
 */
export class BeginRegistrationResponse extends Message<BeginRegistrationResponse> {
  /**
   * @generated from field: string credential_creation_options = 1;
   */
  credentialCreationOptions = "";

  constructor(data?: PartialMessage<BeginRegistrationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.BeginRegistrationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "credential_creation_options", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BeginRegistrationResponse {
    return new BeginRegistrationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BeginRegistrationResponse {
    return new BeginRegistrationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BeginRegistrationResponse {
    return new BeginRegistrationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: BeginRegistrationResponse | PlainMessage<BeginRegistrationResponse> | undefined, b: BeginRegistrationResponse | PlainMessage<BeginRegistrationResponse> | undefined): boolean {
    return proto3.util.equals(BeginRegistrationResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.FinishRegistrationRequest
 */
export class FinishRegistrationRequest extends Message<FinishRegistrationRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  /**
   * @generated from field: string credential_creation_response = 2;
   */
  credentialCreationResponse = "";

  constructor(data?: PartialMessage<FinishRegistrationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.FinishRegistrationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "credential_creation_response", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FinishRegistrationRequest {
    return new FinishRegistrationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FinishRegistrationRequest {
    return new FinishRegistrationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FinishRegistrationRequest {
    return new FinishRegistrationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: FinishRegistrationRequest | PlainMessage<FinishRegistrationRequest> | undefined, b: FinishRegistrationRequest | PlainMessage<FinishRegistrationRequest> | undefined): boolean {
    return proto3.util.equals(FinishRegistrationRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.FinishRegistrationResponse
 */
export class FinishRegistrationResponse extends Message<FinishRegistrationResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  constructor(data?: PartialMessage<FinishRegistrationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.FinishRegistrationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FinishRegistrationResponse {
    return new FinishRegistrationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FinishRegistrationResponse {
    return new FinishRegistrationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FinishRegistrationResponse {
    return new FinishRegistrationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: FinishRegistrationResponse | PlainMessage<FinishRegistrationResponse> | undefined, b: FinishRegistrationResponse | PlainMessage<FinishRegistrationResponse> | undefined): boolean {
    return proto3.util.equals(FinishRegistrationResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.BeginLoginRequest
 */
export class BeginLoginRequest extends Message<BeginLoginRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  constructor(data?: PartialMessage<BeginLoginRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.BeginLoginRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BeginLoginRequest {
    return new BeginLoginRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BeginLoginRequest {
    return new BeginLoginRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BeginLoginRequest {
    return new BeginLoginRequest().fromJsonString(jsonString, options);
  }

  static equals(a: BeginLoginRequest | PlainMessage<BeginLoginRequest> | undefined, b: BeginLoginRequest | PlainMessage<BeginLoginRequest> | undefined): boolean {
    return proto3.util.equals(BeginLoginRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.BeginLoginResponse
 */
export class BeginLoginResponse extends Message<BeginLoginResponse> {
  /**
   * @generated from field: string credential_request_options = 1;
   */
  credentialRequestOptions = "";

  constructor(data?: PartialMessage<BeginLoginResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.BeginLoginResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "credential_request_options", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BeginLoginResponse {
    return new BeginLoginResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BeginLoginResponse {
    return new BeginLoginResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BeginLoginResponse {
    return new BeginLoginResponse().fromJsonString(jsonString, options);
  }

  static equals(a: BeginLoginResponse | PlainMessage<BeginLoginResponse> | undefined, b: BeginLoginResponse | PlainMessage<BeginLoginResponse> | undefined): boolean {
    return proto3.util.equals(BeginLoginResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.FinishLoginRequest
 */
export class FinishLoginRequest extends Message<FinishLoginRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  /**
   * @generated from field: string credential_request_result = 2;
   */
  credentialRequestResult = "";

  constructor(data?: PartialMessage<FinishLoginRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.FinishLoginRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "credential_request_result", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FinishLoginRequest {
    return new FinishLoginRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FinishLoginRequest {
    return new FinishLoginRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FinishLoginRequest {
    return new FinishLoginRequest().fromJsonString(jsonString, options);
  }

  static equals(a: FinishLoginRequest | PlainMessage<FinishLoginRequest> | undefined, b: FinishLoginRequest | PlainMessage<FinishLoginRequest> | undefined): boolean {
    return proto3.util.equals(FinishLoginRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.FinishLoginResponse
 */
export class FinishLoginResponse extends Message<FinishLoginResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  /**
   * @generated from field: google.protobuf.Timestamp expire_at = 2;
   */
  expireAt?: Timestamp;

  constructor(data?: PartialMessage<FinishLoginResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.FinishLoginResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expire_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FinishLoginResponse {
    return new FinishLoginResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FinishLoginResponse {
    return new FinishLoginResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FinishLoginResponse {
    return new FinishLoginResponse().fromJsonString(jsonString, options);
  }

  static equals(a: FinishLoginResponse | PlainMessage<FinishLoginResponse> | undefined, b: FinishLoginResponse | PlainMessage<FinishLoginResponse> | undefined): boolean {
    return proto3.util.equals(FinishLoginResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.ListCredentialsRequest
 */
export class ListCredentialsRequest extends Message<ListCredentialsRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  constructor(data?: PartialMessage<ListCredentialsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.ListCredentialsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListCredentialsRequest {
    return new ListCredentialsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListCredentialsRequest {
    return new ListCredentialsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListCredentialsRequest {
    return new ListCredentialsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListCredentialsRequest | PlainMessage<ListCredentialsRequest> | undefined, b: ListCredentialsRequest | PlainMessage<ListCredentialsRequest> | undefined): boolean {
    return proto3.util.equals(ListCredentialsRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.ListCredentialsResponse
 */
export class ListCredentialsResponse extends Message<ListCredentialsResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  /**
   * @generated from field: repeated dashboard.v1alpha1.Credential credentials = 2;
   */
  credentials: Credential[] = [];

  constructor(data?: PartialMessage<ListCredentialsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.ListCredentialsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "credentials", kind: "message", T: Credential, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListCredentialsResponse {
    return new ListCredentialsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListCredentialsResponse {
    return new ListCredentialsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListCredentialsResponse {
    return new ListCredentialsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListCredentialsResponse | PlainMessage<ListCredentialsResponse> | undefined, b: ListCredentialsResponse | PlainMessage<ListCredentialsResponse> | undefined): boolean {
    return proto3.util.equals(ListCredentialsResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.Credential
 */
export class Credential extends Message<Credential> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string display_name = 2;
   */
  displayName = "";

  /**
   * @generated from field: google.protobuf.Timestamp timestamp = 3;
   */
  timestamp?: Timestamp;

  constructor(data?: PartialMessage<Credential>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.Credential";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "timestamp", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Credential {
    return new Credential().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Credential {
    return new Credential().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Credential {
    return new Credential().fromJsonString(jsonString, options);
  }

  static equals(a: Credential | PlainMessage<Credential> | undefined, b: Credential | PlainMessage<Credential> | undefined): boolean {
    return proto3.util.equals(Credential, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.DeleteCredentialRequest
 */
export class DeleteCredentialRequest extends Message<DeleteCredentialRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  /**
   * @generated from field: string cred_id = 2;
   */
  credId = "";

  constructor(data?: PartialMessage<DeleteCredentialRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.DeleteCredentialRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cred_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteCredentialRequest {
    return new DeleteCredentialRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteCredentialRequest {
    return new DeleteCredentialRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteCredentialRequest {
    return new DeleteCredentialRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteCredentialRequest | PlainMessage<DeleteCredentialRequest> | undefined, b: DeleteCredentialRequest | PlainMessage<DeleteCredentialRequest> | undefined): boolean {
    return proto3.util.equals(DeleteCredentialRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.DeleteCredentialResponse
 */
export class DeleteCredentialResponse extends Message<DeleteCredentialResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  constructor(data?: PartialMessage<DeleteCredentialResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.DeleteCredentialResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteCredentialResponse {
    return new DeleteCredentialResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteCredentialResponse {
    return new DeleteCredentialResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteCredentialResponse {
    return new DeleteCredentialResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteCredentialResponse | PlainMessage<DeleteCredentialResponse> | undefined, b: DeleteCredentialResponse | PlainMessage<DeleteCredentialResponse> | undefined): boolean {
    return proto3.util.equals(DeleteCredentialResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.UpdateCredentialRequest
 */
export class UpdateCredentialRequest extends Message<UpdateCredentialRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  /**
   * @generated from field: string cred_id = 2;
   */
  credId = "";

  /**
   * @generated from field: string cred_display_name = 3;
   */
  credDisplayName = "";

  constructor(data?: PartialMessage<UpdateCredentialRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.UpdateCredentialRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cred_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "cred_display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateCredentialRequest {
    return new UpdateCredentialRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateCredentialRequest {
    return new UpdateCredentialRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateCredentialRequest {
    return new UpdateCredentialRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateCredentialRequest | PlainMessage<UpdateCredentialRequest> | undefined, b: UpdateCredentialRequest | PlainMessage<UpdateCredentialRequest> | undefined): boolean {
    return proto3.util.equals(UpdateCredentialRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.UpdateCredentialResponse
 */
export class UpdateCredentialResponse extends Message<UpdateCredentialResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  constructor(data?: PartialMessage<UpdateCredentialResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "dashboard.v1alpha1.UpdateCredentialResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateCredentialResponse {
    return new UpdateCredentialResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateCredentialResponse {
    return new UpdateCredentialResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateCredentialResponse {
    return new UpdateCredentialResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateCredentialResponse | PlainMessage<UpdateCredentialResponse> | undefined, b: UpdateCredentialResponse | PlainMessage<UpdateCredentialResponse> | undefined): boolean {
    return proto3.util.equals(UpdateCredentialResponse, a, b);
  }
}
