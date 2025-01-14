---
title: Ent + gRPC is Ready for Usage
author: Rotem Tamir
authorURL: "https://github.com/rotemtam"
authorImageURL: "https://s.gravatar.com/avatar/36b3739951a27d2e37251867b7d44b1a?s=80"
authorTwitter: _rtam
---
A few months ago, we announced the experimental support for 
[generating gRPC services from Ent Schema definitions](https://entgo.io/blog/2021/03/18/generating-a-grpc-server-with-ent). The 
implementation was not complete yet but we wanted to get it out the door for the community to experiment with and provide
us with feedback.

Today, after much feedback from the community, we are happy to announce that the [Ent](https://entgo.io) +
[gRPC](https://grpc.io) integration is "Ready for Usage", this means all of the basic features are complete
and we anticipate that most Ent applications can utilize this integration.

What have we added since our initial announcement?
- [Support for "Optional Fields"](https://entgo.io/docs/grpc-optional-fields) - A common issue with Protobufs 
  is that the way that nil values are represented: a zero-valued primitive field isn't encoded into the binary
  representation. This means that applications cannot distinguish between zero and not-set for primitive fields.
  To support this, the Protobuf project supports some 
  "[Well-Known-Types](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf)" 
  called "wrapper types" that wrap the primitive value with a struct. This wasn't previously supported 
  but now when `entproto` generates a Protobuf message definition, it uses these wrapper types to represent
  "Optional" ent fields:
  ```protobuf {15}
  // Code generated by entproto. DO NOT EDIT.
  syntax = "proto3";
  
  package entpb;
  
  import "google/protobuf/wrappers.proto";
  
  message User {
    int32 id = 1;
  
    string name = 2;
  
    string email_address = 3;
  
    google.protobuf.StringValue alias = 4;
  }
  ```

- [Multi-edge support](https://entgo.io/docs/grpc-edges) - when we released the initial version of  
  `protoc-gen-entgrpc`, we only supported generating gRPC service implementations for "Unique" edges
  (i.e reference at most one entity). Since a [recent version](https://github.com/ent/contrib/commit/bf9430fbba45a808bc054144f9711833c76bf05c),
  the plugin supports the generation of gRPC methods to read and write entities with O2M and M2M relationships.
- [Partial responses](https://entgo.io/docs/grpc-edges#retrieving-edge-ids-for-entities) - By default, edge information
  is not returned by the `Get` method of the service. This is done deliberately because the amount of entities related 
  to an entity is unbound.

  To allow the caller of to specify whether or not to return the edge information or not, the generated service adheres
  to [Google AIP-157](https://google.aip.dev/157) (Partial Responses). In short, the `Get<T>Request` message 
  includes an enum named View, this enum allows the caller to control whether or not this information should be retrieved from the database or not.
  
  ```protobuf {6-12}
  message GetUserRequest {
    int32 id = 1;
  
    View view = 2;
  
    enum View {
      VIEW_UNSPECIFIED = 0;
  
      BASIC = 1;
  
      WITH_EDGE_IDS = 2;
    }
  }
  ```

### Getting Started

- To help everyone get started with the Ent + gRPC integration, we have published an official [Ent + gRPC Tutorial](https://entgo.io/docs/grpc-intro) (and a complimentary [GitHub repo](https://github.com/rotemtam/ent-grpc-example)).
- Do you need help getting started with the integration or have some other question? [Join us on Slack](https://entgo.io/docs/slack) or our [Discord server](https://discord.gg/qZmPgTE6RX).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::