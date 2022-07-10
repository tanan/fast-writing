/**
 * @fileoverview gRPC-Web generated client stub for fastwriting
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')

var models_query_pb = require('./models/query_pb.js')

var models_user_pb = require('./models/user_pb.js')

var models_contents_pb = require('./models/contents_pb.js')
const proto = {};
proto.fastwriting = require('./fast-writing_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.fastwriting.UserServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.fastwriting.UserServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.models.UserId,
 *   !proto.models.User>}
 */
const methodDescriptor_UserService_GetUser = new grpc.web.MethodDescriptor(
  '/fastwriting.UserService/GetUser',
  grpc.web.MethodType.UNARY,
  models_user_pb.UserId,
  models_user_pb.User,
  /**
   * @param {!proto.models.UserId} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  models_user_pb.User.deserializeBinary
);


/**
 * @param {!proto.models.UserId} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.models.User)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.models.User>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.UserServiceClient.prototype.getUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.UserService/GetUser',
      request,
      metadata || {},
      methodDescriptor_UserService_GetUser,
      callback);
};


/**
 * @param {!proto.models.UserId} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.models.User>}
 *     Promise that resolves to the response
 */
proto.fastwriting.UserServicePromiseClient.prototype.getUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.UserService/GetUser',
      request,
      metadata || {},
      methodDescriptor_UserService_GetUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.models.User,
 *   !proto.models.UserId>}
 */
const methodDescriptor_UserService_CreateUser = new grpc.web.MethodDescriptor(
  '/fastwriting.UserService/CreateUser',
  grpc.web.MethodType.UNARY,
  models_user_pb.User,
  models_user_pb.UserId,
  /**
   * @param {!proto.models.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  models_user_pb.UserId.deserializeBinary
);


/**
 * @param {!proto.models.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.models.UserId)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.models.UserId>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.UserServiceClient.prototype.createUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.UserService/CreateUser',
      request,
      metadata || {},
      methodDescriptor_UserService_CreateUser,
      callback);
};


/**
 * @param {!proto.models.User} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.models.UserId>}
 *     Promise that resolves to the response
 */
proto.fastwriting.UserServicePromiseClient.prototype.createUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.UserService/CreateUser',
      request,
      metadata || {},
      methodDescriptor_UserService_CreateUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.models.User,
 *   !proto.models.UserId>}
 */
const methodDescriptor_UserService_UpdateUser = new grpc.web.MethodDescriptor(
  '/fastwriting.UserService/UpdateUser',
  grpc.web.MethodType.UNARY,
  models_user_pb.User,
  models_user_pb.UserId,
  /**
   * @param {!proto.models.User} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  models_user_pb.UserId.deserializeBinary
);


/**
 * @param {!proto.models.User} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.models.UserId)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.models.UserId>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.UserServiceClient.prototype.updateUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.UserService/UpdateUser',
      request,
      metadata || {},
      methodDescriptor_UserService_UpdateUser,
      callback);
};


/**
 * @param {!proto.models.User} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.models.UserId>}
 *     Promise that resolves to the response
 */
proto.fastwriting.UserServicePromiseClient.prototype.updateUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.UserService/UpdateUser',
      request,
      metadata || {},
      methodDescriptor_UserService_UpdateUser);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.fastwriting.WritingServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.fastwriting.WritingServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.models.ContentsId,
 *   !proto.models.Contents>}
 */
const methodDescriptor_WritingService_GetContents = new grpc.web.MethodDescriptor(
  '/fastwriting.WritingService/GetContents',
  grpc.web.MethodType.UNARY,
  models_contents_pb.ContentsId,
  models_contents_pb.Contents,
  /**
   * @param {!proto.models.ContentsId} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  models_contents_pb.Contents.deserializeBinary
);


/**
 * @param {!proto.models.ContentsId} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.models.Contents)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.models.Contents>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.WritingServiceClient.prototype.getContents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.WritingService/GetContents',
      request,
      metadata || {},
      methodDescriptor_WritingService_GetContents,
      callback);
};


/**
 * @param {!proto.models.ContentsId} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.models.Contents>}
 *     Promise that resolves to the response
 */
proto.fastwriting.WritingServicePromiseClient.prototype.getContents =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.WritingService/GetContents',
      request,
      metadata || {},
      methodDescriptor_WritingService_GetContents);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.models.ContentsQueryParams,
 *   !proto.models.ContentsList>}
 */
const methodDescriptor_WritingService_GetContentsList = new grpc.web.MethodDescriptor(
  '/fastwriting.WritingService/GetContentsList',
  grpc.web.MethodType.UNARY,
  models_query_pb.ContentsQueryParams,
  models_contents_pb.ContentsList,
  /**
   * @param {!proto.models.ContentsQueryParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  models_contents_pb.ContentsList.deserializeBinary
);


/**
 * @param {!proto.models.ContentsQueryParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.models.ContentsList)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.models.ContentsList>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.WritingServiceClient.prototype.getContentsList =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.WritingService/GetContentsList',
      request,
      metadata || {},
      methodDescriptor_WritingService_GetContentsList,
      callback);
};


/**
 * @param {!proto.models.ContentsQueryParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.models.ContentsList>}
 *     Promise that resolves to the response
 */
proto.fastwriting.WritingServicePromiseClient.prototype.getContentsList =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.WritingService/GetContentsList',
      request,
      metadata || {},
      methodDescriptor_WritingService_GetContentsList);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.models.TagQueryParams,
 *   !proto.models.ContentsList>}
 */
const methodDescriptor_WritingService_GetContentsListByTags = new grpc.web.MethodDescriptor(
  '/fastwriting.WritingService/GetContentsListByTags',
  grpc.web.MethodType.UNARY,
  models_query_pb.TagQueryParams,
  models_contents_pb.ContentsList,
  /**
   * @param {!proto.models.TagQueryParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  models_contents_pb.ContentsList.deserializeBinary
);


/**
 * @param {!proto.models.TagQueryParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.models.ContentsList)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.models.ContentsList>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.WritingServiceClient.prototype.getContentsListByTags =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.WritingService/GetContentsListByTags',
      request,
      metadata || {},
      methodDescriptor_WritingService_GetContentsListByTags,
      callback);
};


/**
 * @param {!proto.models.TagQueryParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.models.ContentsList>}
 *     Promise that resolves to the response
 */
proto.fastwriting.WritingServicePromiseClient.prototype.getContentsListByTags =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.WritingService/GetContentsListByTags',
      request,
      metadata || {},
      methodDescriptor_WritingService_GetContentsListByTags);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.models.UserContentsQueryParams,
 *   !proto.models.ContentsList>}
 */
const methodDescriptor_WritingService_GetUserContentsList = new grpc.web.MethodDescriptor(
  '/fastwriting.WritingService/GetUserContentsList',
  grpc.web.MethodType.UNARY,
  models_query_pb.UserContentsQueryParams,
  models_contents_pb.ContentsList,
  /**
   * @param {!proto.models.UserContentsQueryParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  models_contents_pb.ContentsList.deserializeBinary
);


/**
 * @param {!proto.models.UserContentsQueryParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.models.ContentsList)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.models.ContentsList>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.WritingServiceClient.prototype.getUserContentsList =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.WritingService/GetUserContentsList',
      request,
      metadata || {},
      methodDescriptor_WritingService_GetUserContentsList,
      callback);
};


/**
 * @param {!proto.models.UserContentsQueryParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.models.ContentsList>}
 *     Promise that resolves to the response
 */
proto.fastwriting.WritingServicePromiseClient.prototype.getUserContentsList =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.WritingService/GetUserContentsList',
      request,
      metadata || {},
      methodDescriptor_WritingService_GetUserContentsList);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.fastwriting.CreateContentsRequest,
 *   !proto.fastwriting.CreateContentsResponse>}
 */
const methodDescriptor_WritingService_CreateUserContents = new grpc.web.MethodDescriptor(
  '/fastwriting.WritingService/CreateUserContents',
  grpc.web.MethodType.UNARY,
  proto.fastwriting.CreateContentsRequest,
  proto.fastwriting.CreateContentsResponse,
  /**
   * @param {!proto.fastwriting.CreateContentsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.fastwriting.CreateContentsResponse.deserializeBinary
);


/**
 * @param {!proto.fastwriting.CreateContentsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.fastwriting.CreateContentsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.fastwriting.CreateContentsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.WritingServiceClient.prototype.createUserContents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.WritingService/CreateUserContents',
      request,
      metadata || {},
      methodDescriptor_WritingService_CreateUserContents,
      callback);
};


/**
 * @param {!proto.fastwriting.CreateContentsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.fastwriting.CreateContentsResponse>}
 *     Promise that resolves to the response
 */
proto.fastwriting.WritingServicePromiseClient.prototype.createUserContents =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.WritingService/CreateUserContents',
      request,
      metadata || {},
      methodDescriptor_WritingService_CreateUserContents);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.fastwriting.DeleteContentsRequest,
 *   !proto.fastwriting.DeleteResponse>}
 */
const methodDescriptor_WritingService_DeleteUserContents = new grpc.web.MethodDescriptor(
  '/fastwriting.WritingService/DeleteUserContents',
  grpc.web.MethodType.UNARY,
  proto.fastwriting.DeleteContentsRequest,
  proto.fastwriting.DeleteResponse,
  /**
   * @param {!proto.fastwriting.DeleteContentsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.fastwriting.DeleteResponse.deserializeBinary
);


/**
 * @param {!proto.fastwriting.DeleteContentsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.fastwriting.DeleteResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.fastwriting.DeleteResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.WritingServiceClient.prototype.deleteUserContents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.WritingService/DeleteUserContents',
      request,
      metadata || {},
      methodDescriptor_WritingService_DeleteUserContents,
      callback);
};


/**
 * @param {!proto.fastwriting.DeleteContentsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.fastwriting.DeleteResponse>}
 *     Promise that resolves to the response
 */
proto.fastwriting.WritingServicePromiseClient.prototype.deleteUserContents =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.WritingService/DeleteUserContents',
      request,
      metadata || {},
      methodDescriptor_WritingService_DeleteUserContents);
};


module.exports = proto.fastwriting;

