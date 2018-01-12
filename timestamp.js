/*eslint-disable block-scoped-var, no-redeclare, no-control-regex, no-prototype-builtins*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.timestamp = (function() {

    /**
     * Namespace timestamp.
     * @exports timestamp
     * @namespace
     */
    var timestamp = {};

    timestamp.Timestamp = (function() {

        /**
         * Properties of a Timestamp.
         * @memberof timestamp
         * @interface ITimestamp
         * @property {number|Long|null} [timeUnixMs] Timestamp timeUnixMs
         */

        /**
         * Constructs a new Timestamp.
         * @memberof timestamp
         * @classdesc Represents a Timestamp.
         * @implements ITimestamp
         * @constructor
         * @param {timestamp.ITimestamp=} [properties] Properties to set
         */
        function Timestamp(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Timestamp timeUnixMs.
         * @member {number|Long} timeUnixMs
         * @memberof timestamp.Timestamp
         * @instance
         */
        Timestamp.prototype.timeUnixMs = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Creates a new Timestamp instance using the specified properties.
         * @function create
         * @memberof timestamp.Timestamp
         * @static
         * @param {timestamp.ITimestamp=} [properties] Properties to set
         * @returns {timestamp.Timestamp} Timestamp instance
         */
        Timestamp.create = function create(properties) {
            return new Timestamp(properties);
        };

        /**
         * Encodes the specified Timestamp message. Does not implicitly {@link timestamp.Timestamp.verify|verify} messages.
         * @function encode
         * @memberof timestamp.Timestamp
         * @static
         * @param {timestamp.ITimestamp} message Timestamp message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Timestamp.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.timeUnixMs != null && message.hasOwnProperty("timeUnixMs"))
                writer.uint32(/* id 1, wireType 0 =*/8).uint64(message.timeUnixMs);
            return writer;
        };

        /**
         * Encodes the specified Timestamp message, length delimited. Does not implicitly {@link timestamp.Timestamp.verify|verify} messages.
         * @function encodeDelimited
         * @memberof timestamp.Timestamp
         * @static
         * @param {timestamp.ITimestamp} message Timestamp message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Timestamp.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Timestamp message from the specified reader or buffer.
         * @function decode
         * @memberof timestamp.Timestamp
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {timestamp.Timestamp} Timestamp
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Timestamp.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.timestamp.Timestamp();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.timeUnixMs = reader.uint64();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Timestamp message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof timestamp.Timestamp
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {timestamp.Timestamp} Timestamp
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Timestamp.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Timestamp message.
         * @function verify
         * @memberof timestamp.Timestamp
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Timestamp.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.timeUnixMs != null && message.hasOwnProperty("timeUnixMs"))
                if (!$util.isInteger(message.timeUnixMs) && !(message.timeUnixMs && $util.isInteger(message.timeUnixMs.low) && $util.isInteger(message.timeUnixMs.high)))
                    return "timeUnixMs: integer|Long expected";
            return null;
        };

        /**
         * Creates a Timestamp message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof timestamp.Timestamp
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {timestamp.Timestamp} Timestamp
         */
        Timestamp.fromObject = function fromObject(object) {
            if (object instanceof $root.timestamp.Timestamp)
                return object;
            var message = new $root.timestamp.Timestamp();
            if (object.timeUnixMs != null)
                if ($util.Long)
                    (message.timeUnixMs = $util.Long.fromValue(object.timeUnixMs)).unsigned = true;
                else if (typeof object.timeUnixMs === "string")
                    message.timeUnixMs = parseInt(object.timeUnixMs, 10);
                else if (typeof object.timeUnixMs === "number")
                    message.timeUnixMs = object.timeUnixMs;
                else if (typeof object.timeUnixMs === "object")
                    message.timeUnixMs = new $util.LongBits(object.timeUnixMs.low >>> 0, object.timeUnixMs.high >>> 0).toNumber(true);
            return message;
        };

        /**
         * Creates a plain object from a Timestamp message. Also converts values to other types if specified.
         * @function toObject
         * @memberof timestamp.Timestamp
         * @static
         * @param {timestamp.Timestamp} message Timestamp
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Timestamp.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.timeUnixMs = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.timeUnixMs = options.longs === String ? "0" : 0;
            if (message.timeUnixMs != null && message.hasOwnProperty("timeUnixMs"))
                if (typeof message.timeUnixMs === "number")
                    object.timeUnixMs = options.longs === String ? String(message.timeUnixMs) : message.timeUnixMs;
                else
                    object.timeUnixMs = options.longs === String ? $util.Long.prototype.toString.call(message.timeUnixMs) : options.longs === Number ? new $util.LongBits(message.timeUnixMs.low >>> 0, message.timeUnixMs.high >>> 0).toNumber(true) : message.timeUnixMs;
            return object;
        };

        /**
         * Converts this Timestamp to JSON.
         * @function toJSON
         * @memberof timestamp.Timestamp
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Timestamp.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Timestamp;
    })();

    return timestamp;
})();

module.exports = $root;
