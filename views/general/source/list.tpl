<div class="col-lg-8 col-md-7 col-sm-6" >
    <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
            <ul class="breadcrumb">
                <li><a href="/admin">Admin</a></li>
                <li class="active">Sources</li>
            </ul></div>
        <div class="panel-body">
            {{ if .anything }}
                <p>This is the list with the sources</p>
                <table class="table table-striped table-hover ">              
                    <thead>
                        <tr>
                            <th>Source</th>
                            <th>URL</th>
                            <th>Description</th>
                            <th style="text-align:left">Options</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range $index, $element := .sources}}                    
                        <tr>
                            <td>{{ $element.Name }}</td>
                            <td><a href="{{ $element.Url }}" target="_blank">{{ $element.Url }}</a></td>
                            <td>{{ $element.Description }}</td>
                            <td>
                                <div class="btn-group">
                                    <a href="/" class="btn btn-link dropdown-toggle btn-sm" data-toggle="dropdown"><span class="caret"></span></a>
                                    <ul class="dropdown-menu">
                                        <li><a href="/admin/sources/modify/{{$element.Id}}">Modify</a></li>
                                        <li><a class="deleteSourceButton" data-id="{{$element.Id}}" data-name="{{$element.Name}}" href="/">Delete</a></li>
                                    </ul>
                                </div>
                            </td>
                        </tr>               
                        {{end }}                    
                    </tbody>
                </table> 
            {{ else }}
                There are no sources... :(
                {{ end }}
            </div>
        </div>
    </div>
    <script>
        /* global bootbox */
        $(document).ready(function () {
            $(".deleteSourceButton").click(confirmDelete);
        });

        function confirmDelete(e) {
            e.preventDefault();

            var instance = $(this),
                id = instance.data("id"),
                name = instance.data("name");
            showConfirmationDialog(id, name);
        }

        function showConfirmationDialog(id, name) {
            bootbox.dialog({
                title: "Confirmation",
                message: "The source <b>" + name + "</b> will be permanently deleted and it cannot be recovered. Are you sure?",
                onEscape: true,
                buttons: {
                  cancel: {
                    label: "Cancel",
                    className: "btn-default",
                    callback: function() {
                        this.modal('hide');
                    }
                  },
                  main: {
                    label: "Delete",
                    className: "btn-primary",
                    callback: function() {
                        deleteSource(id);
                    }
                  }
                }
              });
        }
        function deleteSource(id) {
            var newForm = jQuery('<form>', {
                'action': '/admin/sources/delete/' + id,
                'method': "POST"
            });
            newForm.submit();
        }
    </script>