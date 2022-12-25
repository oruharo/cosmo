//
// Cosmo Dashboard API
// Manipulate cosmo dashboard resource API

// @generated by protoc-gen-connect-web v0.5.0 with parameter "target=ts"
// @generated from file dashboard/v1alpha1/user_service.proto (package dashboard.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserRequest, CreateUserResponse, DeleteUserRequest, DeleteUserResponse, GetUserRequest, GetUserResponse, GetUsersResponse, UpdateUserDisplayNameRequest, UpdateUserDisplayNameResponse, UpdateUserPasswordRequest, UpdateUserPasswordResponse, UpdateUserRoleRequest, UpdateUserRoleResponse } from "./user_service_pb.js";
import { Empty, MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service dashboard.v1alpha1.UserService
 */
export const UserService = {
  typeName: "dashboard.v1alpha1.UserService",
  methods: {
    /**
     * Delete user by ID
     *
     * @generated from rpc dashboard.v1alpha1.UserService.DeleteUser
     */
    deleteUser: {
      name: "DeleteUser",
      I: DeleteUserRequest,
      O: DeleteUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Returns a single User model
     *
     * @generated from rpc dashboard.v1alpha1.UserService.GetUser
     */
    getUser: {
      name: "GetUser",
      I: GetUserRequest,
      O: GetUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Returns an array of User model
     *
     * @generated from rpc dashboard.v1alpha1.UserService.GetUsers
     */
    getUsers: {
      name: "GetUsers",
      I: Empty,
      O: GetUsersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Create a new User
     *
     * @generated from rpc dashboard.v1alpha1.UserService.CreateUser
     */
    createUser: {
      name: "CreateUser",
      I: CreateUserRequest,
      O: CreateUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Update user display name
     *
     * @generated from rpc dashboard.v1alpha1.UserService.UpdateUserDisplayName
     */
    updateUserDisplayName: {
      name: "UpdateUserDisplayName",
      I: UpdateUserDisplayNameRequest,
      O: UpdateUserDisplayNameResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Update a single User password
     *
     * @generated from rpc dashboard.v1alpha1.UserService.UpdateUserPassword
     */
    updateUserPassword: {
      name: "UpdateUserPassword",
      I: UpdateUserPasswordRequest,
      O: UpdateUserPasswordResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Update a single User role
     *
     * @generated from rpc dashboard.v1alpha1.UserService.UpdateUserRole
     */
    updateUserRole: {
      name: "UpdateUserRole",
      I: UpdateUserRoleRequest,
      O: UpdateUserRoleResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

