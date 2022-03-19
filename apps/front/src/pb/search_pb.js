// source: search.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
var models_query_pb = require('./models/query_pb.js');
goog.object.extend(proto, models_query_pb);
var models_contents_pb = require('./models/contents_pb.js');
goog.object.extend(proto, models_contents_pb);
goog.exportSymbol('proto.fastwriting.ContentsScore', null, global);
goog.exportSymbol('proto.fastwriting.ContentsScoreList', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fastwriting.ContentsScore = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fastwriting.ContentsScore, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.fastwriting.ContentsScore.displayName = 'proto.fastwriting.ContentsScore';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fastwriting.ContentsScoreList = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.fastwriting.ContentsScoreList.repeatedFields_, null);
};
goog.inherits(proto.fastwriting.ContentsScoreList, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.fastwriting.ContentsScoreList.displayName = 'proto.fastwriting.ContentsScoreList';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fastwriting.ContentsScore.prototype.toObject = function(opt_includeInstance) {
  return proto.fastwriting.ContentsScore.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fastwriting.ContentsScore} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fastwriting.ContentsScore.toObject = function(includeInstance, msg) {
  var f, obj = {
    contentsid: (f = msg.getContentsid()) && models_contents_pb.ContentsId.toObject(includeInstance, f),
    score: jspb.Message.getFieldWithDefault(msg, 2, 0),
    lastUpdated: (f = msg.getLastUpdated()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fastwriting.ContentsScore}
 */
proto.fastwriting.ContentsScore.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fastwriting.ContentsScore;
  return proto.fastwriting.ContentsScore.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fastwriting.ContentsScore} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fastwriting.ContentsScore}
 */
proto.fastwriting.ContentsScore.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new models_contents_pb.ContentsId;
      reader.readMessage(value,models_contents_pb.ContentsId.deserializeBinaryFromReader);
      msg.setContentsid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setScore(value);
      break;
    case 3:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLastUpdated(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fastwriting.ContentsScore.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fastwriting.ContentsScore.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fastwriting.ContentsScore} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fastwriting.ContentsScore.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getContentsid();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      models_contents_pb.ContentsId.serializeBinaryToWriter
    );
  }
  f = message.getScore();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getLastUpdated();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional models.ContentsId contentsId = 1;
 * @return {?proto.models.ContentsId}
 */
proto.fastwriting.ContentsScore.prototype.getContentsid = function() {
  return /** @type{?proto.models.ContentsId} */ (
    jspb.Message.getWrapperField(this, models_contents_pb.ContentsId, 1));
};


/**
 * @param {?proto.models.ContentsId|undefined} value
 * @return {!proto.fastwriting.ContentsScore} returns this
*/
proto.fastwriting.ContentsScore.prototype.setContentsid = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.fastwriting.ContentsScore} returns this
 */
proto.fastwriting.ContentsScore.prototype.clearContentsid = function() {
  return this.setContentsid(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.fastwriting.ContentsScore.prototype.hasContentsid = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int64 score = 2;
 * @return {number}
 */
proto.fastwriting.ContentsScore.prototype.getScore = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.fastwriting.ContentsScore} returns this
 */
proto.fastwriting.ContentsScore.prototype.setScore = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional google.protobuf.Timestamp last_updated = 3;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.fastwriting.ContentsScore.prototype.getLastUpdated = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.fastwriting.ContentsScore} returns this
*/
proto.fastwriting.ContentsScore.prototype.setLastUpdated = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.fastwriting.ContentsScore} returns this
 */
proto.fastwriting.ContentsScore.prototype.clearLastUpdated = function() {
  return this.setLastUpdated(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.fastwriting.ContentsScore.prototype.hasLastUpdated = function() {
  return jspb.Message.getField(this, 3) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.fastwriting.ContentsScoreList.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fastwriting.ContentsScoreList.prototype.toObject = function(opt_includeInstance) {
  return proto.fastwriting.ContentsScoreList.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fastwriting.ContentsScoreList} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fastwriting.ContentsScoreList.toObject = function(includeInstance, msg) {
  var f, obj = {
    contentsscoreList: jspb.Message.toObjectList(msg.getContentsscoreList(),
    proto.fastwriting.ContentsScore.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fastwriting.ContentsScoreList}
 */
proto.fastwriting.ContentsScoreList.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fastwriting.ContentsScoreList;
  return proto.fastwriting.ContentsScoreList.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fastwriting.ContentsScoreList} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fastwriting.ContentsScoreList}
 */
proto.fastwriting.ContentsScoreList.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.fastwriting.ContentsScore;
      reader.readMessage(value,proto.fastwriting.ContentsScore.deserializeBinaryFromReader);
      msg.addContentsscore(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fastwriting.ContentsScoreList.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fastwriting.ContentsScoreList.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fastwriting.ContentsScoreList} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fastwriting.ContentsScoreList.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getContentsscoreList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.fastwriting.ContentsScore.serializeBinaryToWriter
    );
  }
};


/**
 * repeated ContentsScore contentsScore = 1;
 * @return {!Array<!proto.fastwriting.ContentsScore>}
 */
proto.fastwriting.ContentsScoreList.prototype.getContentsscoreList = function() {
  return /** @type{!Array<!proto.fastwriting.ContentsScore>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.fastwriting.ContentsScore, 1));
};


/**
 * @param {!Array<!proto.fastwriting.ContentsScore>} value
 * @return {!proto.fastwriting.ContentsScoreList} returns this
*/
proto.fastwriting.ContentsScoreList.prototype.setContentsscoreList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.fastwriting.ContentsScore=} opt_value
 * @param {number=} opt_index
 * @return {!proto.fastwriting.ContentsScore}
 */
proto.fastwriting.ContentsScoreList.prototype.addContentsscore = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.fastwriting.ContentsScore, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.fastwriting.ContentsScoreList} returns this
 */
proto.fastwriting.ContentsScoreList.prototype.clearContentsscoreList = function() {
  return this.setContentsscoreList([]);
};


goog.object.extend(exports, proto.fastwriting);
