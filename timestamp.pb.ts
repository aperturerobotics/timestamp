/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal.js";

export const protobufPackage = "timestamp";

/** Timestamp contains a cross-platform timestamp. */
export interface Timestamp {
  /** TimeUnixMs timestamp in unix milliseconds. */
  timeUnixMs: Long;
}

function createBaseTimestamp(): Timestamp {
  return { timeUnixMs: Long.UZERO };
}

export const Timestamp = {
  encode(message: Timestamp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (!message.timeUnixMs.isZero()) {
      writer.uint32(8).uint64(message.timeUnixMs);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Timestamp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTimestamp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.timeUnixMs = reader.uint64() as Long;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  // encodeTransform encodes a source of message objects.
  // Transform<Timestamp, Uint8Array>
  async *encodeTransform(
    source: AsyncIterable<Timestamp | Timestamp[]> | Iterable<Timestamp | Timestamp[]>,
  ): AsyncIterable<Uint8Array> {
    for await (const pkt of source) {
      if (Array.isArray(pkt)) {
        for (const p of pkt) {
          yield* [Timestamp.encode(p).finish()];
        }
      } else {
        yield* [Timestamp.encode(pkt).finish()];
      }
    }
  },

  // decodeTransform decodes a source of encoded messages.
  // Transform<Uint8Array, Timestamp>
  async *decodeTransform(
    source: AsyncIterable<Uint8Array | Uint8Array[]> | Iterable<Uint8Array | Uint8Array[]>,
  ): AsyncIterable<Timestamp> {
    for await (const pkt of source) {
      if (Array.isArray(pkt)) {
        for (const p of pkt) {
          yield* [Timestamp.decode(p)];
        }
      } else {
        yield* [Timestamp.decode(pkt)];
      }
    }
  },

  fromJSON(object: any): Timestamp {
    return { timeUnixMs: isSet(object.timeUnixMs) ? Long.fromValue(object.timeUnixMs) : Long.UZERO };
  },

  toJSON(message: Timestamp): unknown {
    const obj: any = {};
    message.timeUnixMs !== undefined && (obj.timeUnixMs = (message.timeUnixMs || Long.UZERO).toString());
    return obj;
  },

  create<I extends Exact<DeepPartial<Timestamp>, I>>(base?: I): Timestamp {
    return Timestamp.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Timestamp>, I>>(object: I): Timestamp {
    const message = createBaseTimestamp();
    message.timeUnixMs = (object.timeUnixMs !== undefined && object.timeUnixMs !== null)
      ? Long.fromValue(object.timeUnixMs)
      : Long.UZERO;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Long ? string | number | Long : T extends Array<infer U> ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends { $case: string } ? { [K in keyof Omit<T, "$case">]?: DeepPartial<T[K]> } & { $case: T["$case"] }
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
