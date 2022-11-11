//
// Cosmo Dashboard API
// Manipulate cosmo dashboard resource API

// @generated by protoc-gen-es v0.2.1 with parameter "target=ts"
// @generated from file dashboard/v1alpha1/user_service.proto (package dashboard.v1alpha1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";
import {User, UserAddons} from "./user_pb.js";

/**
 * @generated from message dashboard.v1alpha1.DeleteUserRequest
 */
export class DeleteUserRequest extends Message<DeleteUserRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  constructor(data?: PartialMessage<DeleteUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.DeleteUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserRequest {
    return new DeleteUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserRequest {
    return new DeleteUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserRequest {
    return new DeleteUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteUserRequest | PlainMessage<DeleteUserRequest> | undefined, b: DeleteUserRequest | PlainMessage<DeleteUserRequest> | undefined): boolean {
    return proto3.util.equals(DeleteUserRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.DeleteUserResponse
 */
export class DeleteUserResponse extends Message<DeleteUserResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  /**
   * @generated from field: dashboard.v1alpha1.User user = 2;
   */
  user?: User;

  constructor(data?: PartialMessage<DeleteUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.DeleteUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserResponse {
    return new DeleteUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserResponse {
    return new DeleteUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserResponse {
    return new DeleteUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteUserResponse | PlainMessage<DeleteUserResponse> | undefined, b: DeleteUserResponse | PlainMessage<DeleteUserResponse> | undefined): boolean {
    return proto3.util.equals(DeleteUserResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.GetUsersResponse
 */
export class GetUsersResponse extends Message<GetUsersResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  /**
   * @generated from field: repeated dashboard.v1alpha1.User items = 2;
   */
  items: User[] = [];

  constructor(data?: PartialMessage<GetUsersResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.GetUsersResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "items", kind: "message", T: User, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUsersResponse {
    return new GetUsersResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUsersResponse {
    return new GetUsersResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUsersResponse {
    return new GetUsersResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetUsersResponse | PlainMessage<GetUsersResponse> | undefined, b: GetUsersResponse | PlainMessage<GetUsersResponse> | undefined): boolean {
    return proto3.util.equals(GetUsersResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.GetUserRequest
 */
export class GetUserRequest extends Message<GetUserRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  constructor(data?: PartialMessage<GetUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.GetUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserRequest {
    return new GetUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserRequest {
    return new GetUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserRequest {
    return new GetUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserRequest | PlainMessage<GetUserRequest> | undefined, b: GetUserRequest | PlainMessage<GetUserRequest> | undefined): boolean {
    return proto3.util.equals(GetUserRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.GetUserResponse
 */
export class GetUserResponse extends Message<GetUserResponse> {
  /**
   * @generated from field: dashboard.v1alpha1.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<GetUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.GetUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserResponse {
    return new GetUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserResponse {
    return new GetUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserResponse {
    return new GetUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserResponse | PlainMessage<GetUserResponse> | undefined, b: GetUserResponse | PlainMessage<GetUserResponse> | undefined): boolean {
    return proto3.util.equals(GetUserResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.CreateUserRequest
 */
export class CreateUserRequest extends Message<CreateUserRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  /**
   * @generated from field: string display_name = 2;
   */
  displayName = "";

  /**
   * @generated from field: string role = 3;
   */
  role = "";

  /**
   * @generated from field: string auth_type = 4;
   */
  authType = "";

  /**
   * @generated from field: repeated dashboard.v1alpha1.UserAddons addons = 5;
   */
  addons: UserAddons[] = [];

  constructor(data?: PartialMessage<CreateUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.CreateUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "role", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "auth_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "addons", kind: "message", T: UserAddons, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined, b: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined): boolean {
    return proto3.util.equals(CreateUserRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.CreateUserResponse
 */
export class CreateUserResponse extends Message<CreateUserResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  /**
   * @generated from field: dashboard.v1alpha1.User user = 2;
   */
  user?: User;

  constructor(data?: PartialMessage<CreateUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.CreateUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined, b: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined): boolean {
    return proto3.util.equals(CreateUserResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.UpdateUserDisplayNameRequest
 */
export class UpdateUserDisplayNameRequest extends Message<UpdateUserDisplayNameRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  /**
   * @generated from field: string display_name = 2;
   */
  displayName = "";

  constructor(data?: PartialMessage<UpdateUserDisplayNameRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.UpdateUserDisplayNameRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserDisplayNameRequest {
    return new UpdateUserDisplayNameRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserDisplayNameRequest {
    return new UpdateUserDisplayNameRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserDisplayNameRequest {
    return new UpdateUserDisplayNameRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserDisplayNameRequest | PlainMessage<UpdateUserDisplayNameRequest> | undefined, b: UpdateUserDisplayNameRequest | PlainMessage<UpdateUserDisplayNameRequest> | undefined): boolean {
    return proto3.util.equals(UpdateUserDisplayNameRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.UpdateUserDisplayNameResponse
 */
export class UpdateUserDisplayNameResponse extends Message<UpdateUserDisplayNameResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  /**
   * @generated from field: dashboard.v1alpha1.User user = 2;
   */
  user?: User;

  constructor(data?: PartialMessage<UpdateUserDisplayNameResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.UpdateUserDisplayNameResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserDisplayNameResponse {
    return new UpdateUserDisplayNameResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserDisplayNameResponse {
    return new UpdateUserDisplayNameResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserDisplayNameResponse {
    return new UpdateUserDisplayNameResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserDisplayNameResponse | PlainMessage<UpdateUserDisplayNameResponse> | undefined, b: UpdateUserDisplayNameResponse | PlainMessage<UpdateUserDisplayNameResponse> | undefined): boolean {
    return proto3.util.equals(UpdateUserDisplayNameResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.UpdateUserPasswordRequest
 */
export class UpdateUserPasswordRequest extends Message<UpdateUserPasswordRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  /**
   * @generated from field: string current_password = 2;
   */
  currentPassword = "";

  /**
   * @generated from field: string new_password = 3;
   */
  newPassword = "";

  constructor(data?: PartialMessage<UpdateUserPasswordRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.UpdateUserPasswordRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "current_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "new_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserPasswordRequest {
    return new UpdateUserPasswordRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserPasswordRequest {
    return new UpdateUserPasswordRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserPasswordRequest {
    return new UpdateUserPasswordRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserPasswordRequest | PlainMessage<UpdateUserPasswordRequest> | undefined, b: UpdateUserPasswordRequest | PlainMessage<UpdateUserPasswordRequest> | undefined): boolean {
    return proto3.util.equals(UpdateUserPasswordRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.UpdateUserPasswordResponse
 */
export class UpdateUserPasswordResponse extends Message<UpdateUserPasswordResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  constructor(data?: PartialMessage<UpdateUserPasswordResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.UpdateUserPasswordResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserPasswordResponse {
    return new UpdateUserPasswordResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserPasswordResponse {
    return new UpdateUserPasswordResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserPasswordResponse {
    return new UpdateUserPasswordResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserPasswordResponse | PlainMessage<UpdateUserPasswordResponse> | undefined, b: UpdateUserPasswordResponse | PlainMessage<UpdateUserPasswordResponse> | undefined): boolean {
    return proto3.util.equals(UpdateUserPasswordResponse, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.UpdateUserRoleRequest
 */
export class UpdateUserRoleRequest extends Message<UpdateUserRoleRequest> {
  /**
   * @generated from field: string user_name = 1;
   */
  userName = "";

  /**
   * @generated from field: string role = 2;
   */
  role = "";

  constructor(data?: PartialMessage<UpdateUserRoleRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.UpdateUserRoleRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "role", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserRoleRequest {
    return new UpdateUserRoleRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserRoleRequest {
    return new UpdateUserRoleRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserRoleRequest {
    return new UpdateUserRoleRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserRoleRequest | PlainMessage<UpdateUserRoleRequest> | undefined, b: UpdateUserRoleRequest | PlainMessage<UpdateUserRoleRequest> | undefined): boolean {
    return proto3.util.equals(UpdateUserRoleRequest, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.UpdateUserRoleResponse
 */
export class UpdateUserRoleResponse extends Message<UpdateUserRoleResponse> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  /**
   * @generated from field: dashboard.v1alpha1.User user = 2;
   */
  user?: User;

  constructor(data?: PartialMessage<UpdateUserRoleResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.UpdateUserRoleResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserRoleResponse {
    return new UpdateUserRoleResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserRoleResponse {
    return new UpdateUserRoleResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserRoleResponse {
    return new UpdateUserRoleResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserRoleResponse | PlainMessage<UpdateUserRoleResponse> | undefined, b: UpdateUserRoleResponse | PlainMessage<UpdateUserRoleResponse> | undefined): boolean {
    return proto3.util.equals(UpdateUserRoleResponse, a, b);
  }
}

