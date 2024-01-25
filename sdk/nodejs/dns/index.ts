// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetRecordArgs, GetRecordResult, GetRecordOutputArgs } from "./getRecord";
export const getRecord: typeof import("./getRecord").getRecord = null as any;
export const getRecordOutput: typeof import("./getRecord").getRecordOutput = null as any;
utilities.lazyLoad(exports, ["getRecord","getRecordOutput"], () => require("./getRecord"));

export { GetZoneArgs, GetZoneResult, GetZoneOutputArgs } from "./getZone";
export const getZone: typeof import("./getZone").getZone = null as any;
export const getZoneOutput: typeof import("./getZone").getZoneOutput = null as any;
utilities.lazyLoad(exports, ["getZone","getZoneOutput"], () => require("./getZone"));

export { RecordArgs, RecordState } from "./record";
export type Record = import("./record").Record;
export const Record: typeof import("./record").Record = null as any;
utilities.lazyLoad(exports, ["Record"], () => require("./record"));

export { ZoneArgs, ZoneState } from "./zone";
export type Zone = import("./zone").Zone;
export const Zone: typeof import("./zone").Zone = null as any;
utilities.lazyLoad(exports, ["Zone"], () => require("./zone"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway:dns/record:Record":
                return new Record(name, <any>undefined, { urn })
            case "scaleway:dns/zone:Zone":
                return new Zone(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway", "dns/record", _module)
pulumi.runtime.registerResourceModule("scaleway", "dns/zone", _module)
