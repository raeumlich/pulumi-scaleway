// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { APIKeyArgs, APIKeyState } from "./apikey";
export type APIKey = import("./apikey").APIKey;
export const APIKey: typeof import("./apikey").APIKey = null as any;
utilities.lazyLoad(exports, ["APIKey"], () => require("./apikey"));

export { ApplicationArgs, ApplicationState } from "./application";
export type Application = import("./application").Application;
export const Application: typeof import("./application").Application = null as any;
utilities.lazyLoad(exports, ["Application"], () => require("./application"));

export { GetApplicationArgs, GetApplicationResult, GetApplicationOutputArgs } from "./getApplication";
export const getApplication: typeof import("./getApplication").getApplication = null as any;
export const getApplicationOutput: typeof import("./getApplication").getApplicationOutput = null as any;
utilities.lazyLoad(exports, ["getApplication","getApplicationOutput"], () => require("./getApplication"));

export { GetGroupArgs, GetGroupResult, GetGroupOutputArgs } from "./getGroup";
export const getGroup: typeof import("./getGroup").getGroup = null as any;
export const getGroupOutput: typeof import("./getGroup").getGroupOutput = null as any;
utilities.lazyLoad(exports, ["getGroup","getGroupOutput"], () => require("./getGroup"));

export { GetSSHKeyArgs, GetSSHKeyResult, GetSSHKeyOutputArgs } from "./getSSHKey";
export const getSSHKey: typeof import("./getSSHKey").getSSHKey = null as any;
export const getSSHKeyOutput: typeof import("./getSSHKey").getSSHKeyOutput = null as any;
utilities.lazyLoad(exports, ["getSSHKey","getSSHKeyOutput"], () => require("./getSSHKey"));

export { GetUserArgs, GetUserResult, GetUserOutputArgs } from "./getUser";
export const getUser: typeof import("./getUser").getUser = null as any;
export const getUserOutput: typeof import("./getUser").getUserOutput = null as any;
utilities.lazyLoad(exports, ["getUser","getUserOutput"], () => require("./getUser"));

export { GroupArgs, GroupState } from "./group";
export type Group = import("./group").Group;
export const Group: typeof import("./group").Group = null as any;
utilities.lazyLoad(exports, ["Group"], () => require("./group"));

export { GroupMembershipArgs, GroupMembershipState } from "./groupMembership";
export type GroupMembership = import("./groupMembership").GroupMembership;
export const GroupMembership: typeof import("./groupMembership").GroupMembership = null as any;
utilities.lazyLoad(exports, ["GroupMembership"], () => require("./groupMembership"));

export { PolicyArgs, PolicyState } from "./policy";
export type Policy = import("./policy").Policy;
export const Policy: typeof import("./policy").Policy = null as any;
utilities.lazyLoad(exports, ["Policy"], () => require("./policy"));

export { SSHKeyArgs, SSHKeyState } from "./sshkey";
export type SSHKey = import("./sshkey").SSHKey;
export const SSHKey: typeof import("./sshkey").SSHKey = null as any;
utilities.lazyLoad(exports, ["SSHKey"], () => require("./sshkey"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway:iam/aPIKey:APIKey":
                return new APIKey(name, <any>undefined, { urn })
            case "scaleway:iam/application:Application":
                return new Application(name, <any>undefined, { urn })
            case "scaleway:iam/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "scaleway:iam/groupMembership:GroupMembership":
                return new GroupMembership(name, <any>undefined, { urn })
            case "scaleway:iam/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "scaleway:iam/sSHKey:SSHKey":
                return new SSHKey(name, <any>undefined, { urn })
            case "scaleway:iam/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway", "iam/aPIKey", _module)
pulumi.runtime.registerResourceModule("scaleway", "iam/application", _module)
pulumi.runtime.registerResourceModule("scaleway", "iam/group", _module)
pulumi.runtime.registerResourceModule("scaleway", "iam/groupMembership", _module)
pulumi.runtime.registerResourceModule("scaleway", "iam/policy", _module)
pulumi.runtime.registerResourceModule("scaleway", "iam/sSHKey", _module)
pulumi.runtime.registerResourceModule("scaleway", "iam/user", _module)