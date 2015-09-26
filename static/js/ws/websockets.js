/* global $, wisply*/
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
     * @memberof Websockets
     * @class Connection
     * @classdesc It represents a websocket connection
     * @param {string} host The host address
     * @param {object} info An object which contians the callbacks for the connection (open, errror, message)
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
    /** @lends Websockets.Connection */
    {
        /**
         * It sends a message
         * @param  {object} message The object to be send
         */
        send: function (message) {
            this.value.send(JSON.stringify(message));
        }
    };

    /**
     * Gets the JQuery element and show waiting
     * @memberof Websockets
     * @class Gui
     * @classdesc The GUI is used by the Connection to show the progress.
     */
    var GUI = function GUI() {
        this.element = $("#websocket-connection");
        this.showWaiting();
    };
    GUI.prototype =
        /** @lends Websockets.GUI */
        {
            /**
             * It shows the wisply waiting message
             */
            showWaiting: function() {
              this.setText(wisply.getLoadingImage("small"));
            },
            /**
             * It shows that there is no connection
             */
            showError: function() {
              this.setText("<span class='text-danger'>No live connection <span class='glyphicon glyphicon-adjust'></span></span>");
            },
            /**
             * It shows the conneciton is live
             */
            showSuccess: function() {
              this.setText("<span class='text-success'>Live connection <span class='glyphicon glyphicon-adjust'></span></span>");
            },
            /**
             * It changes the html of the element
             * @param  {string} text The HTML to be inserted
             */
            setText :function(text) {
              this.element.html(text);
            }
        };
    return {
        Connection: Connection
    };
};
