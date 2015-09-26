/* Global $,wisply*/
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
        var copyInfo = info,
        instance = this;
        this.gui = new GUI();
        this.value = new WebSocket("ws://" + host);

        var open = (function () {
            var conn = instance,
              i = copyInfo;
            return function () {
              conn.gui.showSuccess();
              i.onOpen();
            };
        })();

        var error = (function () {
          var conn = instance,
            i = copyInfo;
            return function () {
              conn.gui.showError();
              i.onError();
            };
        })();


        var message = (function () {
          var i = copyInfo;
            return function (evt) {
              i.onMessage(evt);
            };
        })();

        this.value.onopen = open;
        this.value.onclose = error;
        this.value.onerror = error;
        this.value.onmessage = message;
        this.status = "wait";
    };
    Connection.prototype =
    /** @lends Harvest.Connection */
    {
        send: function (message) {
            this.value.send(JSON.stringify(message));
        }
    };


    var GUI = function GUI() {
        this.element = $("#websocket-connection");
        this.showWaiting();
    };
    GUI.prototype =
        /** @lends ListHarvest.GUI */
        {
            showWaiting: function() {
              this.setText(wisply.getLoadingImage("small"));
            },
            showError: function() {
              this.setText("<span class='text-danger'>No live connection <span class='glyphicon glyphicon-adjust'></span></span>");
            },
            showSuccess: function() {
              this.setText("<span class='text-success'>Live connection <span class='glyphicon glyphicon-adjust'></span></span>");
            },
            setText :function(text) {
              this.element.html(text);
            }
        };

    return {
        Connection: Connection
    };
};
