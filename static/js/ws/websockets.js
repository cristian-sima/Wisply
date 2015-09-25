/**
 * @file Encapsulates the functionality for connecting using web sockets
 * @author Cristian Sima
 */
/**
 * @namespace Websockets
 */
var Websockets = function () {
    'use strict';
    /**
     * Starts the connection and inits the listeners
     * @memberof Harvest
     * @class Connection
     * @classdesc It represents a websocket connection
     * @param {object} info
     */
    var Connection = function Connection(host, info) {
        this.value = new WebSocket("ws://" + host);
        this.value.onopen = info.onOpen;
        this.value.onclose = info.onClose;
        this.value.onerror = info.onError;
        this.status = "wait";
    };
    Connection.prototype =
    /** @lends Harvest.Connection */
    {
        send: function (message) {
            this.value.send(JSON.stringify(message));
        }
    };
    return {
        Connection: Connection
    };
};
