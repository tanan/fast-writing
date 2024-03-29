// source: models/contents.proto
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
var global = (function() { return this || window || global || self || Function('return this')(); }).call(null);

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.models.Contents', null, global);
goog.exportSymbol('proto.models.ContentsId', null, global);
goog.exportSymbol('proto.models.ContentsList', null, global);
goog.exportSymbol('proto.models.Quiz', null, global);
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
proto.models.ContentsId = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.models.ContentsId, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.models.ContentsId.displayName = 'proto.models.ContentsId';
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
proto.models.Contents = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.models.Contents.repeatedFields_, null);
};
goog.inherits(proto.models.Contents, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.models.Contents.displayName = 'proto.models.Contents';
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
proto.models.ContentsList = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.models.ContentsList.repeatedFields_, null);
};
goog.inherits(proto.models.ContentsList, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.models.ContentsList.displayName = 'proto.models.ContentsList';
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
proto.models.Quiz = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.models.Quiz, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.models.Quiz.displayName = 'proto.models.Quiz';
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
proto.models.ContentsId.prototype.toObject = function(opt_includeInstance) {
  return proto.models.ContentsId.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.models.ContentsId} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.models.ContentsId.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, 0)
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
 * @return {!proto.models.ContentsId}
 */
proto.models.ContentsId.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.models.ContentsId;
  return proto.models.ContentsId.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.models.ContentsId} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.models.ContentsId}
 */
proto.models.ContentsId.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setId(value);
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
proto.models.ContentsId.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.models.ContentsId.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.models.ContentsId} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.models.ContentsId.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
};


/**
 * optional int64 id = 1;
 * @return {number}
 */
proto.models.ContentsId.prototype.getId = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.models.ContentsId} returns this
 */
proto.models.ContentsId.prototype.setId = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.models.Contents.repeatedFields_ = [6];



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
proto.models.Contents.prototype.toObject = function(opt_includeInstance) {
  return proto.models.Contents.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.models.Contents} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.models.Contents.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: (f = msg.getId()) && proto.models.ContentsId.toObject(includeInstance, f),
    title: jspb.Message.getFieldWithDefault(msg, 2, ""),
    description: jspb.Message.getFieldWithDefault(msg, 3, ""),
    creator: jspb.Message.getFieldWithDefault(msg, 4, ""),
    scope: jspb.Message.getFieldWithDefault(msg, 5, ""),
    quizlistList: jspb.Message.toObjectList(msg.getQuizlistList(),
    proto.models.Quiz.toObject, includeInstance),
    lastUpdated: (f = msg.getLastUpdated()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    tags: jspb.Message.getFieldWithDefault(msg, 8, "")
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
 * @return {!proto.models.Contents}
 */
proto.models.Contents.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.models.Contents;
  return proto.models.Contents.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.models.Contents} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.models.Contents}
 */
proto.models.Contents.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.models.ContentsId;
      reader.readMessage(value,proto.models.ContentsId.deserializeBinaryFromReader);
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTitle(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setDescription(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setCreator(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setScope(value);
      break;
    case 6:
      var value = new proto.models.Quiz;
      reader.readMessage(value,proto.models.Quiz.deserializeBinaryFromReader);
      msg.addQuizlist(value);
      break;
    case 7:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLastUpdated(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setTags(value);
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
proto.models.Contents.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.models.Contents.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.models.Contents} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.models.Contents.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.models.ContentsId.serializeBinaryToWriter
    );
  }
  f = message.getTitle();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDescription();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getCreator();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getScope();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getQuizlistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      6,
      f,
      proto.models.Quiz.serializeBinaryToWriter
    );
  }
  f = message.getLastUpdated();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getTags();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
};


/**
 * optional ContentsId id = 1;
 * @return {?proto.models.ContentsId}
 */
proto.models.Contents.prototype.getId = function() {
  return /** @type{?proto.models.ContentsId} */ (
    jspb.Message.getWrapperField(this, proto.models.ContentsId, 1));
};


/**
 * @param {?proto.models.ContentsId|undefined} value
 * @return {!proto.models.Contents} returns this
*/
proto.models.Contents.prototype.setId = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.models.Contents} returns this
 */
proto.models.Contents.prototype.clearId = function() {
  return this.setId(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.models.Contents.prototype.hasId = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string title = 2;
 * @return {string}
 */
proto.models.Contents.prototype.getTitle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.models.Contents} returns this
 */
proto.models.Contents.prototype.setTitle = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string description = 3;
 * @return {string}
 */
proto.models.Contents.prototype.getDescription = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.models.Contents} returns this
 */
proto.models.Contents.prototype.setDescription = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string creator = 4;
 * @return {string}
 */
proto.models.Contents.prototype.getCreator = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.models.Contents} returns this
 */
proto.models.Contents.prototype.setCreator = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string scope = 5;
 * @return {string}
 */
proto.models.Contents.prototype.getScope = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.models.Contents} returns this
 */
proto.models.Contents.prototype.setScope = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * repeated Quiz quizList = 6;
 * @return {!Array<!proto.models.Quiz>}
 */
proto.models.Contents.prototype.getQuizlistList = function() {
  return /** @type{!Array<!proto.models.Quiz>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.models.Quiz, 6));
};


/**
 * @param {!Array<!proto.models.Quiz>} value
 * @return {!proto.models.Contents} returns this
*/
proto.models.Contents.prototype.setQuizlistList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 6, value);
};


/**
 * @param {!proto.models.Quiz=} opt_value
 * @param {number=} opt_index
 * @return {!proto.models.Quiz}
 */
proto.models.Contents.prototype.addQuizlist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 6, opt_value, proto.models.Quiz, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.models.Contents} returns this
 */
proto.models.Contents.prototype.clearQuizlistList = function() {
  return this.setQuizlistList([]);
};


/**
 * optional google.protobuf.Timestamp last_updated = 7;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.models.Contents.prototype.getLastUpdated = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 7));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.models.Contents} returns this
*/
proto.models.Contents.prototype.setLastUpdated = function(value) {
  return jspb.Message.setWrapperField(this, 7, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.models.Contents} returns this
 */
proto.models.Contents.prototype.clearLastUpdated = function() {
  return this.setLastUpdated(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.models.Contents.prototype.hasLastUpdated = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional string tags = 8;
 * @return {string}
 */
proto.models.Contents.prototype.getTags = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.models.Contents} returns this
 */
proto.models.Contents.prototype.setTags = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.models.ContentsList.repeatedFields_ = [1];



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
proto.models.ContentsList.prototype.toObject = function(opt_includeInstance) {
  return proto.models.ContentsList.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.models.ContentsList} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.models.ContentsList.toObject = function(includeInstance, msg) {
  var f, obj = {
    contentslistList: jspb.Message.toObjectList(msg.getContentslistList(),
    proto.models.Contents.toObject, includeInstance)
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
 * @return {!proto.models.ContentsList}
 */
proto.models.ContentsList.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.models.ContentsList;
  return proto.models.ContentsList.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.models.ContentsList} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.models.ContentsList}
 */
proto.models.ContentsList.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.models.Contents;
      reader.readMessage(value,proto.models.Contents.deserializeBinaryFromReader);
      msg.addContentslist(value);
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
proto.models.ContentsList.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.models.ContentsList.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.models.ContentsList} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.models.ContentsList.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getContentslistList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.models.Contents.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Contents contentsList = 1;
 * @return {!Array<!proto.models.Contents>}
 */
proto.models.ContentsList.prototype.getContentslistList = function() {
  return /** @type{!Array<!proto.models.Contents>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.models.Contents, 1));
};


/**
 * @param {!Array<!proto.models.Contents>} value
 * @return {!proto.models.ContentsList} returns this
*/
proto.models.ContentsList.prototype.setContentslistList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.models.Contents=} opt_value
 * @param {number=} opt_index
 * @return {!proto.models.Contents}
 */
proto.models.ContentsList.prototype.addContentslist = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.models.Contents, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.models.ContentsList} returns this
 */
proto.models.ContentsList.prototype.clearContentslistList = function() {
  return this.setContentslistList([]);
};





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
proto.models.Quiz.prototype.toObject = function(opt_includeInstance) {
  return proto.models.Quiz.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.models.Quiz} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.models.Quiz.toObject = function(includeInstance, msg) {
  var f, obj = {
    question: jspb.Message.getFieldWithDefault(msg, 1, ""),
    answer: jspb.Message.getFieldWithDefault(msg, 2, ""),
    order: jspb.Message.getFieldWithDefault(msg, 3, 0)
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
 * @return {!proto.models.Quiz}
 */
proto.models.Quiz.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.models.Quiz;
  return proto.models.Quiz.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.models.Quiz} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.models.Quiz}
 */
proto.models.Quiz.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setQuestion(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAnswer(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setOrder(value);
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
proto.models.Quiz.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.models.Quiz.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.models.Quiz} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.models.Quiz.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getQuestion();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAnswer();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getOrder();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
};


/**
 * optional string question = 1;
 * @return {string}
 */
proto.models.Quiz.prototype.getQuestion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.models.Quiz} returns this
 */
proto.models.Quiz.prototype.setQuestion = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string answer = 2;
 * @return {string}
 */
proto.models.Quiz.prototype.getAnswer = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.models.Quiz} returns this
 */
proto.models.Quiz.prototype.setAnswer = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional int64 order = 3;
 * @return {number}
 */
proto.models.Quiz.prototype.getOrder = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.models.Quiz} returns this
 */
proto.models.Quiz.prototype.setOrder = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


goog.object.extend(exports, proto.models);
