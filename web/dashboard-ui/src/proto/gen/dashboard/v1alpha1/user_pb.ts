//
// Cosmo Dashboard API
// Manipulate cosmo dashboard resource API

// @generated by protoc-gen-es v0.2.1 with parameter "target=ts"
// @generated from file dashboard/v1alpha1/user.proto (package dashboard.v1alpha1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";

/**
 * @generated from message dashboard.v1alpha1.User
 */
export class User extends Message<User> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string display_name = 2;
   */
  displayName = "";

  /**
   * {"","cosmo-admin"}
   *
   * @generated from field: string role = 3;
   */
  role = "";

  /**
   * {"","kosmo-secret"}
   *
   * @generated from field: string auth_type = 4;
   */
  authType = "";

  /**
   * @generated from field: repeated dashboard.v1alpha1.UserAddons addons = 5;
   */
  addons: UserAddons[] = [];

  /**
   * @generated from field: string default_password = 6;
   */
  defaultPassword = "";

  /**
   * @generated from field: string status = 7;
   */
  status = "";

  constructor(data?: PartialMessage<User>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.User";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "role", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "auth_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "addons", kind: "message", T: UserAddons, repeated: true },
    { no: 6, name: "default_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "status", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): User {
    return new User().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): User {
    return new User().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): User {
    return new User().fromJsonString(jsonString, options);
  }

  static equals(a: User | PlainMessage<User> | undefined, b: User | PlainMessage<User> | undefined): boolean {
    return proto3.util.equals(User, a, b);
  }
}

/**
 * @generated from message dashboard.v1alpha1.UserAddons
 */
export class UserAddons extends Message<UserAddons> {
  /**
   * @generated from field: string template = 1;
   */
  template = "";

  /**
   * @generated from field: bool cluster_scoped = 2;
   */
  clusterScoped = false;

  /**
   * @generated from field: map<string, string> vars = 3;
   */
  vars: { [key: string]: string } = {};

  constructor(data?: PartialMessage<UserAddons>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "dashboard.v1alpha1.UserAddons";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "template", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cluster_scoped", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "vars", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserAddons {
    return new UserAddons().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserAddons {
    return new UserAddons().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserAddons {
    return new UserAddons().fromJsonString(jsonString, options);
  }

  static equals(a: UserAddons | PlainMessage<UserAddons> | undefined, b: UserAddons | PlainMessage<UserAddons> | undefined): boolean {
    return proto3.util.equals(UserAddons, a, b);
  }
}

