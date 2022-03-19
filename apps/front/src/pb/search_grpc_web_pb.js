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


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var models_query_pb = require('./models/query_pb.js')

var models_contents_pb = require('./models/contents_pb.js')
const proto = {};
proto.fastwriting = require('./search_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.fastwriting.SearchServiceClient =
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
proto.fastwriting.SearchServicePromiseClient =
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
 *   !proto.models.TitleQueryParams,
 *   !proto.fastwriting.ContentsScoreList>}
 */
const methodDescriptor_SearchService_FindContentsIdListByTitle = new grpc.web.MethodDescriptor(
  '/fastwriting.SearchService/FindContentsIdListByTitle',
  grpc.web.MethodType.UNARY,
  models_query_pb.TitleQueryParams,
  proto.fastwriting.ContentsScoreList,
  /**
   * @param {!proto.models.TitleQueryParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.fastwriting.ContentsScoreList.deserializeBinary
);


/**
 * @param {!proto.models.TitleQueryParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.fastwriting.ContentsScoreList)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.fastwriting.ContentsScoreList>|undefined}
 *     The XHR Node Readable Stream
 */
proto.fastwriting.SearchServiceClient.prototype.findContentsIdListByTitle =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/fastwriting.SearchService/FindContentsIdListByTitle',
      request,
      metadata || {},
      methodDescriptor_SearchService_FindContentsIdListByTitle,
      callback);
};


/**
 * @param {!proto.models.TitleQueryParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.fastwriting.ContentsScoreList>}
 *     Promise that resolves to the response
 */
proto.fastwriting.SearchServicePromiseClient.prototype.findContentsIdListByTitle =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/fastwriting.SearchService/FindContentsIdListByTitle',
      request,
      metadata || {},
      methodDescriptor_SearchService_FindContentsIdListByTitle);
};


module.exports = proto.fastwriting;

