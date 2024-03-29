// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BucketArgs, BucketState } from "./bucket";
export type Bucket = import("./bucket").Bucket;
export const Bucket: typeof import("./bucket").Bucket = null as any;
utilities.lazyLoad(exports, ["Bucket"], () => require("./bucket"));

export { BucketACLArgs, BucketACLState } from "./bucketACL";
export type BucketACL = import("./bucketACL").BucketACL;
export const BucketACL: typeof import("./bucketACL").BucketACL = null as any;
utilities.lazyLoad(exports, ["BucketACL"], () => require("./bucketACL"));

export { BucketLockConfigurationArgs, BucketLockConfigurationState } from "./bucketLockConfiguration";
export type BucketLockConfiguration = import("./bucketLockConfiguration").BucketLockConfiguration;
export const BucketLockConfiguration: typeof import("./bucketLockConfiguration").BucketLockConfiguration = null as any;
utilities.lazyLoad(exports, ["BucketLockConfiguration"], () => require("./bucketLockConfiguration"));

export { BucketPolicyArgs, BucketPolicyState } from "./bucketPolicy";
export type BucketPolicy = import("./bucketPolicy").BucketPolicy;
export const BucketPolicy: typeof import("./bucketPolicy").BucketPolicy = null as any;
utilities.lazyLoad(exports, ["BucketPolicy"], () => require("./bucketPolicy"));

export { BucketWebsiteConfigurationArgs, BucketWebsiteConfigurationState } from "./bucketWebsiteConfiguration";
export type BucketWebsiteConfiguration = import("./bucketWebsiteConfiguration").BucketWebsiteConfiguration;
export const BucketWebsiteConfiguration: typeof import("./bucketWebsiteConfiguration").BucketWebsiteConfiguration = null as any;
utilities.lazyLoad(exports, ["BucketWebsiteConfiguration"], () => require("./bucketWebsiteConfiguration"));

export { GetBucketArgs, GetBucketResult, GetBucketOutputArgs } from "./getBucket";
export const getBucket: typeof import("./getBucket").getBucket = null as any;
export const getBucketOutput: typeof import("./getBucket").getBucketOutput = null as any;
utilities.lazyLoad(exports, ["getBucket","getBucketOutput"], () => require("./getBucket"));

export { GetBucketPolicyArgs, GetBucketPolicyResult, GetBucketPolicyOutputArgs } from "./getBucketPolicy";
export const getBucketPolicy: typeof import("./getBucketPolicy").getBucketPolicy = null as any;
export const getBucketPolicyOutput: typeof import("./getBucketPolicy").getBucketPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getBucketPolicy","getBucketPolicyOutput"], () => require("./getBucketPolicy"));

export { ItemArgs, ItemState } from "./item";
export type Item = import("./item").Item;
export const Item: typeof import("./item").Item = null as any;
utilities.lazyLoad(exports, ["Item"], () => require("./item"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "scaleway:objectstorage/bucket:Bucket":
                return new Bucket(name, <any>undefined, { urn })
            case "scaleway:objectstorage/bucketACL:BucketACL":
                return new BucketACL(name, <any>undefined, { urn })
            case "scaleway:objectstorage/bucketLockConfiguration:BucketLockConfiguration":
                return new BucketLockConfiguration(name, <any>undefined, { urn })
            case "scaleway:objectstorage/bucketPolicy:BucketPolicy":
                return new BucketPolicy(name, <any>undefined, { urn })
            case "scaleway:objectstorage/bucketWebsiteConfiguration:BucketWebsiteConfiguration":
                return new BucketWebsiteConfiguration(name, <any>undefined, { urn })
            case "scaleway:objectstorage/item:Item":
                return new Item(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("scaleway", "objectstorage/bucket", _module)
pulumi.runtime.registerResourceModule("scaleway", "objectstorage/bucketACL", _module)
pulumi.runtime.registerResourceModule("scaleway", "objectstorage/bucketLockConfiguration", _module)
pulumi.runtime.registerResourceModule("scaleway", "objectstorage/bucketPolicy", _module)
pulumi.runtime.registerResourceModule("scaleway", "objectstorage/bucketWebsiteConfiguration", _module)
pulumi.runtime.registerResourceModule("scaleway", "objectstorage/item", _module)
