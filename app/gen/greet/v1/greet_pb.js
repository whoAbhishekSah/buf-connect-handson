// @generated by protoc-gen-es v1.3.0
// @generated from file greet/v1/greet.proto (package greet.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message greet.v1.GreetRequest
 */
export const GreetRequest = proto3.makeMessageType(
  "greet.v1.GreetRequest",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message greet.v1.GreetResponse
 */
export const GreetResponse = proto3.makeMessageType(
  "greet.v1.GreetResponse",
  () => [
    { no: 1, name: "greeting", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "greeting_number", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

