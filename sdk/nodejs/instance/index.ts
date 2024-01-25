// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetIPArgs, GetIPResult, GetIPOutputArgs } from "./getIP";
export const getIP: typeof import("./getIP").getIP = null as any;
export const getIPOutput: typeof import("./getIP").getIPOutput = null as any;
utilities.lazyLoad(exports, ["getIP","getIPOutput"], () => require("./getIP"));

export { GetImageArgs, GetImageResult, GetImageOutputArgs } from "./getImage";
export const getImage: typeof import("./getImage").getImage = null as any;
export const getImageOutput: typeof import("./getImage").getImageOutput = null as any;
utilities.lazyLoad(exports, ["getImage","getImageOutput"], () => require("./getImage"));

export { GetMarketplaceImageArgs, GetMarketplaceImageResult, GetMarketplaceImageOutputArgs } from "./getMarketplaceImage";
export const getMarketplaceImage: typeof import("./getMarketplaceImage").getMarketplaceImage = null as any;
export const getMarketplaceImageOutput: typeof import("./getMarketplaceImage").getMarketplaceImageOutput = null as any;
utilities.lazyLoad(exports, ["getMarketplaceImage","getMarketplaceImageOutput"], () => require("./getMarketplaceImage"));

export { GetPlacementGroupArgs, GetPlacementGroupResult, GetPlacementGroupOutputArgs } from "./getPlacementGroup";
export const getPlacementGroup: typeof import("./getPlacementGroup").getPlacementGroup = null as any;
export const getPlacementGroupOutput: typeof import("./getPlacementGroup").getPlacementGroupOutput = null as any;
utilities.lazyLoad(exports, ["getPlacementGroup","getPlacementGroupOutput"], () => require("./getPlacementGroup"));

export { GetPrivateNICArgs, GetPrivateNICResult, GetPrivateNICOutputArgs } from "./getPrivateNIC";
export const getPrivateNIC: typeof import("./getPrivateNIC").getPrivateNIC = null as any;
export const getPrivateNICOutput: typeof import("./getPrivateNIC").getPrivateNICOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateNIC","getPrivateNICOutput"], () => require("./getPrivateNIC"));

export { GetSecurityGroupArgs, GetSecurityGroupResult, GetSecurityGroupOutputArgs } from "./getSecurityGroup";
export const getSecurityGroup: typeof import("./getSecurityGroup").getSecurityGroup = null as any;
export const getSecurityGroupOutput: typeof import("./getSecurityGroup").getSecurityGroupOutput = null as any;
utilities.lazyLoad(exports, ["getSecurityGroup","getSecurityGroupOutput"], () => require("./getSecurityGroup"));

export { GetServerArgs, GetServerResult, GetServerOutputArgs } from "./getServer";
export const getServer: typeof import("./getServer").getServer = null as any;
export const getServerOutput: typeof import("./getServer").getServerOutput = null as any;
utilities.lazyLoad(exports, ["getServer","getServerOutput"], () => require("./getServer"));

export { GetServersArgs, GetServersResult, GetServersOutputArgs } from "./getServers";
export const getServers: typeof import("./getServers").getServers = null as any;
export const getServersOutput: typeof import("./getServers").getServersOutput = null as any;
utilities.lazyLoad(exports, ["getServers","getServersOutput"], () => require("./getServers"));

export { GetSnapshotArgs, GetSnapshotResult, GetSnapshotOutputArgs } from "./getSnapshot";
export const getSnapshot: typeof import("./getSnapshot").getSnapshot = null as any;
export const getSnapshotOutput: typeof import("./getSnapshot").getSnapshotOutput = null as any;
utilities.lazyLoad(exports, ["getSnapshot","getSnapshotOutput"], () => require("./getSnapshot"));

export { GetVolumeArgs, GetVolumeResult, GetVolumeOutputArgs } from "./getVolume";
export const getVolume: typeof import("./getVolume").getVolume = null as any;
export const getVolumeOutput: typeof import("./getVolume").getVolumeOutput = null as any;
utilities.lazyLoad(exports, ["getVolume","getVolumeOutput"], () => require("./getVolume"));

export { ImageArgs, ImageState } from "./image";
export type Image = import("./image").Image;
export const Image: typeof import("./image").Image = null as any;
utilities.lazyLoad(exports, ["Image"], () => require("./image"));

export { IPArgs, IPState } from "./ip";
export type IP = import("./ip").IP;
export const IP: typeof import("./ip").IP = null as any;
utilities.lazyLoad(exports, ["IP"], () => require("./ip"));

export { IPReverseDNSArgs, IPReverseDNSState } from "./ipreverseDNS";
export type IPReverseDNS = import("./ipreverseDNS").IPReverseDNS;
export const IPReverseDNS: typeof import("./ipreverseDNS").IPReverseDNS = null as any;
utilities.lazyLoad(exports, ["IPReverseDNS"], () => require("./ipreverseDNS"));

export { PlacementGroupArgs, PlacementGroupState } from "./placementGroup";
export type PlacementGroup = import("./placementGroup").PlacementGroup;
export const PlacementGroup: typeof import("./placementGroup").PlacementGroup = null as any;
utilities.lazyLoad(exports, ["PlacementGroup"], () => require("./placementGroup"));

export { PrivateNICArgs, PrivateNICState } from "./privateNIC";
export type PrivateNIC = import("./privateNIC").PrivateNIC;
export const PrivateNIC: typeof import("./privateNIC").PrivateNIC = null as any;
utilities.lazyLoad(exports, ["PrivateNIC"], () => require("./privateNIC"));

export { SecurityGroupArgs, SecurityGroupState } from "./securityGroup";
export type SecurityGroup = import("./securityGroup").SecurityGroup;
export const SecurityGroup: typeof import("./securityGroup").SecurityGroup = null as any;
utilities.lazyLoad(exports, ["SecurityGroup"], () => require("./securityGroup"));

export { SecurityGroupRulesArgs, SecurityGroupRulesState } from "./securityGroupRules";
export type SecurityGroupRules = import("./securityGroupRules").SecurityGroupRules;
export const SecurityGroupRules: typeof import("./securityGroupRules").SecurityGroupRules = null as any;
utilities.lazyLoad(exports, ["SecurityGroupRules"], () => require("./securityGroupRules"));

export { ServerArgs, ServerState } from "./server";
export type Server = import("./server").Server;
export const Server: typeof import("./server").Server = null as any;
utilities.lazyLoad(exports, ["Server"], () => require("./server"));

export { SnapshotArgs, SnapshotState } from "./snapshot";
export type Snapshot = import("./snapshot").Snapshot;
export const Snapshot: typeof import("./snapshot").Snapshot = null as any;
utilities.lazyLoad(exports, ["Snapshot"], () => require("./snapshot"));

export { UserDataArgs, UserDataState } from "./userData";
export type UserData = import("./userData").UserData;
export const UserData: typeof import("./userData").UserData = null as any;
utilities.lazyLoad(exports, ["UserData"], () => require("./userData"));

export { VolumeArgs, VolumeState } from "./volume";
export type Volume = import("./volume").Volume;
export const Volume: typeof import("./volume").Volume = null as any;
utilities.lazyLoad(exports, ["Volume"], () => require("./volume"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway:instance/iP:IP":
                return new IP(name, <any>undefined, { urn })
            case "scaleway:instance/iPReverseDNS:IPReverseDNS":
                return new IPReverseDNS(name, <any>undefined, { urn })
            case "scaleway:instance/image:Image":
                return new Image(name, <any>undefined, { urn })
            case "scaleway:instance/placementGroup:PlacementGroup":
                return new PlacementGroup(name, <any>undefined, { urn })
            case "scaleway:instance/privateNIC:PrivateNIC":
                return new PrivateNIC(name, <any>undefined, { urn })
            case "scaleway:instance/securityGroup:SecurityGroup":
                return new SecurityGroup(name, <any>undefined, { urn })
            case "scaleway:instance/securityGroupRules:SecurityGroupRules":
                return new SecurityGroupRules(name, <any>undefined, { urn })
            case "scaleway:instance/server:Server":
                return new Server(name, <any>undefined, { urn })
            case "scaleway:instance/snapshot:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "scaleway:instance/userData:UserData":
                return new UserData(name, <any>undefined, { urn })
            case "scaleway:instance/volume:Volume":
                return new Volume(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway", "instance/iP", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/iPReverseDNS", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/image", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/placementGroup", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/privateNIC", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/securityGroup", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/securityGroupRules", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/server", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/snapshot", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/userData", _module)
pulumi.runtime.registerResourceModule("scaleway", "instance/volume", _module)
