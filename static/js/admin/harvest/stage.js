
    /**
     * Saves the stages
     * @memberof Harvest
     * @class StageManager
     * @classdesc It encapsulets the functionality for the sources
     * @param [Manager] repositoriesManager The reference to the repositories manager
     */
    var StageManager = function StageManager(repositoriesManager) {
        this.status = "stopped";
        this.repo = repositoriesManager;
        this.current = 0;
        this.stage = {};
        // stages
        this.data = [{
            name: "Prepare resources",
            id: 0,
            /**
             * It prepares the resources
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                this.paint();
                var manager = stageManager;
                setTimeout(function () {
                    manager.firedStageFinished();
                }, 500);
            },
            /**
             * It shows the loading image
             */
            paint: function () {
                $('#current').html(wisply.getLoadingImage("big"));
            }
        }, {
            name: "Connect to server",
            id: 1,
            /**
             * It tries to connect to server using websockets and saves the connection
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                var manager = stageManager.repo,
                    repository = manager.repository,
                    intance = this,
                    connection = {},
                    copyReferenceManager = stageManager;
                if (window.WebSocket) {
                    connection = {
                        process: function (data) {
                            wisply.harvest.processMessage(data);
                        },
                        open: function() {
                          setTimeout(function () {
                            instance.decide(copyReferenceManager);
                          }, 400);
                        },
                        host: data.host
                    };
                    manager.connection = new Connection(connection);
                } else {
                    this.complain();
                    return;
                }
            },
            /**
             * It tells the user the connection was not done
             */
            complain: function () {
                $('#current').html("Wisply was not able to realize the connection. Your browser does not support WebSockets");
            },
            /**
             * It shows the loading image
             */
            paint: function () {
                $('#current').html(wisply.getLoadingImage("big"));
            }
        },
        {
             name: "Getting current state...",
             id: 2,
             /**
              * It gets the current state from the server
              * @param  {Manager} stageManager The reference to the repositories manager
              */
             perform: function (stageManager) {
                instance.repo.connection.sendMessage("getCurrentProcess", "");
             },
            decide: function (stageManager) {
                if(stageManager.server.hasProcess) {
                    stageManager.repo.history.log("A process already exists for the current repository. The current action on the server  is " + stageManager.currentStages[stageManager.server.CurrentOperation]);
                    stageManager.current = stageManager.server.CurrentOperation;
                }
                stageManager.firedStageFinished();
                // processOnServer


                   if (resultFromServer.IsValid === true) {
                       stageManager.repo.history.log("The URL is valid");
                       stageManager.firedStageFinished();
                   } else {
                       this.complain(stageManager);
                       this.enableModifyURL();
                   }
               },
             /**
              * It shows the loading image
              */
             paint: function () {
                 $('#current').html(wisply.getLoadingImage("big"));
             },
             /**
              * It tells the user that the URL was not valid. It allows the user to edit it
              * @param  {Manager} stageManager The reference to the repositories manager
              */
             complain: function (stageManager) {
                 $('#current').html("The URL is not valid or the address can not be visited. Please correct it and click 'Modify'");
                 stageManager.repo.history.log("The URL is not valid!");
                 stageManager.repo.pause();
             },
             /**
              * It disables the possibility to modify the URL
              */
             disableModifyURL: function () {
                 $('#modifyButton').prop('disabled', true);
                 $('#Source-URL').prop('disabled', true);
             },
             /**
              * It enables the possibility to modify the URL
              */
             enableModifyURL: function () {
                 $('#modifyButton').prop('disabled', false);
                 $('#Source-URL').prop('disabled', false);
             }
         },
         {
            name: "Verify URL address",
            id: 3,
            showBar: false,
            /**
             * It tries to verify the address
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                this.paint();
                this.disableModifyURL();
                var instance = stageManager;
                instance.repo.history.log("Verify URL address");
                instance.repo.connection.sendMessage("testURL", "");
            },
            /**
             * It checks if the server has verified the URL
             * @param  {Manager} stageManager The reference to the repositories manager
             * @param  {object} content      The content of the message from server
             */
            result: function (stageManager, resultFromServer) {
                if (resultFromServer.IsValid === true) {
                    stageManager.repo.history.log("The URL is valid");
                    stageManager.firedStageFinished();
                } else {
                    this.complain(stageManager);
                    this.enableModifyURL();
                }
            },
            /**
             * It shows the loading image
             */
            paint: function () {
                $('#current').html(wisply.getLoadingImage("big"));
            },
            /**
             * It tells the user that the URL was not valid. It allows the user to edit it
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            complain: function (stageManager) {
                $('#current').html("The URL is not valid or the address can not be visited. Please correct it and click 'Modify'");
                stageManager.repo.history.log("The URL is not valid!");
                stageManager.repo.pause();
            },
            /**
             * It disables the possibility to modify the URL
             */
            disableModifyURL: function () {
                $('#modifyButton').prop('disabled', true);
                $('#Source-URL').prop('disabled', true);
            },
            /**
             * It enables the possibility to modify the URL
             */
            enableModifyURL: function () {
                $('#modifyButton').prop('disabled', false);
                $('#Source-URL').prop('disabled', false);
            }
        }, {
            name: "Identify Source",
            id: 4,
            /**
             * It tells the server to identify the source
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                var instance = stageManager;
                this.disableModifyURL();
                instance.repo.history.log("Indentifing the source");
                instance.repo.connection.sendMessage("identify", "something");
            },
            /**
             * It checks if the server has identified the repository
             * @param  {object} indentifyInfo The value of the message from the server
             */
            result: function (stageManager, indentifyInfo) {
                if (indentifyInfo.state === true) {
                    this.paint(indentifyInfo.data.Identify);
                    stageManager.repo.history.log("The source has been identified");
                    stageManager.firedStageFinished();
                } else {
                    this.complain(stageManager);
                    this.enableModifyURL();
                }
            },
            /**
             * It tells the user that the identification has not been done
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            complain: function (stageManager) {
                $('#current').html("There were problems with the indentification");
                stageManager.repo.history.log("There has error during identification!");
                stageManager.repo.pause();
            },
            /**
             * It shows in the current table the details about the repository
             */
            paint: function (data) {
                /**
                 * It returns the HTML table for an object
                 * @param  {object} data The object
                 * @return {string} HTML table for the object's data
                 */
                function getIdentification(data) {
                    var html = "";
                    html += '<ul class="list-group text-left">';
                    for (var property in data) {
                        if (data.hasOwnProperty(property)) {
                            if (property === "Description") {
                                continue;
                            } else if (typeof data[property] === 'object') {
                                html += '<li class="list-group-item"> ' + property;
                                html += getIdentification(data[property]);
                                html += '</li>';
                            } else {
                                html += "  <li class='list-group-item'>";
                                html += property + ": <strong>" + data[property] + "</strong>";
                                html += "</li>";
                            }
                        }
                    }
                    html += "</ul>";
                    return html;
                }
                function drawBadges() {
                  var text = '<div class="row text-center" id="repository-elements" style="display:none"><table class="table"><tbody><tr>';
                  var records = '<td><div class="text-center col-xs-3 col-md-3">' +
                      'Records<br />' +
                        '<span class="badge" id="repository-elements-records">0</span>' +
                     '</div></td>';
                     text += records;
                     text += "</tr></tbody></table></div>";
                  return text;
                }
                var html = "";
                html += drawBadges();
                html += getIdentification(data);
                $("#current").html(html);
            },
            /**
             * It disables the possibility to modify the URL
             */
            disableModifyURL: function () {
                $('#modifyButton').prop('disabled', true);
                $('#Source-URL').prop('disabled', true);
            },
            /**
             * It enables the possibility to modify the URL
             */
            enableModifyURL: function () {
                $('#modifyButton').prop('disabled', false);
                $('#Source-URL').prop('disabled', false);
            }
        },
        {
            name: "Initialize",
            id: 5,
            /**
             * It tells the server to receive the records
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                this.init();
                var instance = stageManager;
                instance.repo.history.log("Telling the server to start the init process...");
                instance.repo.connection.sendMessage("initialize", "");
            },
            /**
             * It checks if the server has identified the repository
             * @param  {object} indentifyInfo The value of the message from the server
             */
            result: function (stageManager, result) {
                $("#repository-elements").slideUp();
                stageManager.firedStageFinished();
            },
            /**
             * It hides the URL and button for modify and shows the name of the repository
             */
            init: function () {
                $("#URL-input").toggle();
                $("#Name-Repository").toggle();
                $("#modifyButton").hide();
            }
        },
        {
            name: "Receiving records...",
            id: 6,
            /**
             * It tells the server to receive the records
             * @param  {Manager} stageManager The reference to the repositories manager
             */
            perform: function (stageManager) {
                var instance = stageManager;
                this.element = $("#repository-elements-records");
                this.start();
                instance.repo.history.log("Prepare to receive records");
            },
            /**
             * It checks if the server has identified the repository
             * @param  {object} indentifyInfo The value of the message from the server
             */
            result: function (stageManager) {
                stageManager.firedStageFinished();
                this.end();
            },
            start: function () {
              $("#repository-elements").slideDown();
              this.element.addClass("progress-bar-warning");
            },
            end: function (stageManager) {
              this.element.removeClass("progress-bar-warning");
              this.element.addClass("progress-bar-success");
              this.element = null;
              stageManager.firedStageFinished();
            },
            update: function (newValue) {
              this.element.html(newValue.Number);
            }
        },
      ];
    };
    StageManager.prototype =
        /** @lends Harvest.StageManager */
        {
            loadConfiguration: function(server) {
              var process;
                if(server.hasProcess) {
                    process = server.currentProcesses;
                }
            },
            /**
             * It starts the manager. It calls the first stage
             */
            start: function () {
                this.performStage(this.current);
            },
            /**
             * It calls the next stage. If there are no stages, it calls firedEnd
             */
            next: function () {
                this.current++;
                if (this.current >= this.data.length) {
                    this.firedEnd();
                } else {
                    this.performStage(this.current);
                }
            },
            /**
             * It performs a stage
             * @param  {number} id The id of the stage
             */
            performStage: function (id) {
                var stage = this.data[id];
                this.status = "running";
                this.repo.history.log("Starting stage <b>" + stage.name + "</b>...");
                this.repo.page.update();
                this.stage = stage;
                stage.perform(this);
            },
            /**
             * It forces the manager to stop. It forces the current stage to stop
             */
            forceStop: function () {
                this.repo.history.log("The stage manager has been forced to stop.");
                this.current = "Stopped";
                this.state = "stopped";
                if (this.stage.stop) {
                    this.stage.stop();
                }
            },
            /**
             * Called when a stage has finished. It updates the page and calls the next stage
             */
            firedStageFinished: function () {
                if (this.state === "stopped" || this.state === "paused") {
                    this.repo.history.log("Imposible to continue!");
                } else {
                    this.repo.page.update();
                    this.repo.history.log("Stage " + (this.current + 1) + " finished!");
                    this.next();
                }
            },
            /**
             * It is called when all the stages has been called. It updates the page
             */
            firedEnd: function () {
                this.status = "finish";
                this.repo.page.update();
                this.repo.history.log("The process has been finished!");
            },
            /**
             * It performs again a stage
             * @param  {number} number The id of the stage
             */
            restart: function (number) {
                this.repo.history.log("Restarting from stage " + (number + 1) + "...");
                this.current = number - 1;
                this.next();
            },
            /**
             * It pauses the manager
             */
            pause: function () {
                this.status = "paused";
            }
        };
