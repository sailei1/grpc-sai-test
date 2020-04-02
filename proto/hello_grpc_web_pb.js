/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('./google/api/annotations_pb.js')
const proto = {};
proto.proto = require('./hello_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.HelloWorldClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

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
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.HelloWorldPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

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
 *   !proto.proto.HelloWorldRequest,
 *   !proto.proto.HelloWorldResponse>}
 */
const methodDescriptor_HelloWorld_SayHelloWorld = new grpc.web.MethodDescriptor(
  '/proto.HelloWorld/SayHelloWorld',
  grpc.web.MethodType.UNARY,
  proto.proto.HelloWorldRequest,
  proto.proto.HelloWorldResponse,
  /**
   * @param {!proto.proto.HelloWorldRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.HelloWorldResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.HelloWorldRequest,
 *   !proto.proto.HelloWorldResponse>}
 */
const methodInfo_HelloWorld_SayHelloWorld = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.HelloWorldResponse,
  /**
   * @param {!proto.proto.HelloWorldRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.HelloWorldResponse.deserializeBinary
);


/**
 * @param {!proto.proto.HelloWorldRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.HelloWorldResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.HelloWorldResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.HelloWorldClient.prototype.sayHelloWorld =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.HelloWorld/SayHelloWorld',
      request,
      metadata || {},
      methodDescriptor_HelloWorld_SayHelloWorld,
      callback);
};


/**
 * @param {!proto.proto.HelloWorldRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.HelloWorldResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.HelloWorldPromiseClient.prototype.sayHelloWorld =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.HelloWorld/SayHelloWorld',
      request,
      metadata || {},
      methodDescriptor_HelloWorld_SayHelloWorld);
};


module.exports = proto.proto;

