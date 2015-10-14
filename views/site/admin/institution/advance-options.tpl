<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/institutions">Institutions</a></li>
      <li><a href="/admin/institutions/institution/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
      <li class="active">Advance options</li>
    </ul></div>
    <div class="panel-body">
      <div class="row" >
        <div class="col-md-12">
          <div>
            <h2>Modify</h2>
            <div class="well">
              You can use this option to modify the details of the institution. <br/>
              <a href="/admin/institutions/modify/{{ .institution.ID }}" class="btn btn-primary">Modify institution</a>
            </div>
          </div>
          <div>
            <h2>Delete</h2>
            <div class="well">
              <h5><span class="text-warning"><span class="glyphicon glyphicon-warning-sign"></span> Please note that the institution can not be recovered.</span></h5><br />
              <h2 class="text-warning">Warning:</h2>
              <div>These will be deleted:</div>
              <ul>
                <li> The information about the institution</li>
                <li> All the repositories
                  <ul>
                    <li>All the records</li>
                    <li>All the collections</li>
                    <li>All the identifiers</li>
                    <li>All the formats</li>
                    <li>All the logs</li>
                  </ul>
                </li>
              </ul>
              <a data-id="" data-name="" href="#" class="btn btn-danger deleteInstitutionButton">Delete entire information regarding institution from Wisply</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script src="/static/js/admin/institution/list.js"></script>
