// source: errorcode.proto
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

goog.exportSymbol('proto.ErrorCode', null, global);
/**
 * @enum {number}
 */
proto.ErrorCode = {
  SUCCESS: 0,
  GATEWAYEXCEPTION: 1,
  NOFINDSERVICE: 10,
  NOFINDSERVICEHANDLEFUNC: 11,
  RPCFUNCEXECUTIONERROR: 12,
  CACHEREADERROR: 13,
  SQLEXECUTIONERROR: 14,
  REQPARAMETERERROR: 15,
  SIGNERROR: 16,
  INSUFFICIENTPERMISSIONS: 17,
  NOLOGIN: 18,
  USERSESSIONNOBEING: 19,
  STATEINVALID: 20,
  DBERROR: 21,
  SYSTEMERROR: 22,
  DECODEERROR: 23,
  TIMESTAMPTIMEOUT: 24,
  PBERROR: 25,
  AGENTUIDEMPTY: 26,
  EXCEPTION: 100,
  AGENTIDUNREGISTERED: 1001,
  AGENTSIGNERROR: 1002
};

goog.object.extend(exports, proto);
