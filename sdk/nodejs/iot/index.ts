// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DeviceArgs, DeviceState } from "./device";
export type Device = import("./device").Device;
export const Device: typeof import("./device").Device = null as any;
utilities.lazyLoad(exports, ["Device"], () => require("./device"));

export { GetDeviceArgs, GetDeviceResult, GetDeviceOutputArgs } from "./getDevice";
export const getDevice: typeof import("./getDevice").getDevice = null as any;
export const getDeviceOutput: typeof import("./getDevice").getDeviceOutput = null as any;
utilities.lazyLoad(exports, ["getDevice","getDeviceOutput"], () => require("./getDevice"));

export { GetHubArgs, GetHubResult, GetHubOutputArgs } from "./getHub";
export const getHub: typeof import("./getHub").getHub = null as any;
export const getHubOutput: typeof import("./getHub").getHubOutput = null as any;
utilities.lazyLoad(exports, ["getHub","getHubOutput"], () => require("./getHub"));

export { HubArgs, HubState } from "./hub";
export type Hub = import("./hub").Hub;
export const Hub: typeof import("./hub").Hub = null as any;
utilities.lazyLoad(exports, ["Hub"], () => require("./hub"));

export { NetworkArgs, NetworkState } from "./network";
export type Network = import("./network").Network;
export const Network: typeof import("./network").Network = null as any;
utilities.lazyLoad(exports, ["Network"], () => require("./network"));

export { RouteArgs, RouteState } from "./route";
export type Route = import("./route").Route;
export const Route: typeof import("./route").Route = null as any;
utilities.lazyLoad(exports, ["Route"], () => require("./route"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway:iot/device:Device":
                return new Device(name, <any>undefined, { urn })
            case "scaleway:iot/hub:Hub":
                return new Hub(name, <any>undefined, { urn })
            case "scaleway:iot/network:Network":
                return new Network(name, <any>undefined, { urn })
            case "scaleway:iot/route:Route":
                return new Route(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway", "iot/device", _module)
pulumi.runtime.registerResourceModule("scaleway", "iot/hub", _module)
pulumi.runtime.registerResourceModule("scaleway", "iot/network", _module)
pulumi.runtime.registerResourceModule("scaleway", "iot/route", _module)
