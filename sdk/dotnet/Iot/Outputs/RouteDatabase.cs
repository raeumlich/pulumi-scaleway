// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Scaleway.Iot.Outputs
{

    [OutputType]
    public sealed class RouteDatabase
    {
        /// <summary>
        /// The database name (e.g. `measurements`).
        /// </summary>
        public readonly string Dbname;
        /// <summary>
        /// The database hostname. Can be an IP or a FQDN.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The database password.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The database port (e.g. `5432`)
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// The SQL query that will be executed when receiving a message ($TOPIC and $PAYLOAD variables are available, see documentation, e.g. `INSERT INTO mytable(date, topic, value) VALUES (NOW(), $TOPIC, $PAYLOAD)`).
        /// </summary>
        public readonly string Query;
        /// <summary>
        /// The database username.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private RouteDatabase(
            string dbname,

            string host,

            string password,

            int port,

            string query,

            string username)
        {
            Dbname = dbname;
            Host = host;
            Password = password;
            Port = port;
            Query = query;
            Username = username;
        }
    }
}