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
                        <th>Name</th>
                        <th>URL</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .sources}}
                    <tr>
                        <td>{{.}}</td>
                        <td><a href="{{.}}">{{.}}<a/></td>
                    </tr>
                    {{end}}                    
                </tbody>
            </table> 
            {{ else }}
             There are no sources... :(
            {{ end }}
        </div>
    </div>
</div>
