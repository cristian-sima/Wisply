<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/education">Education</a></li>
      <li><a href="/admin/education/subjects/{{ .subject.GetID }}">{{ .subject.GetName }}</a></li>
      <li class="active">Advance options</li>
    </ul></div>
    <div class="panel-body">
      <div class="row" >
        <div class="col-md-12">
          <div>
            <h2>Static description</h2>
            <div class="well">
              The static description appears on the public page. <br/>
              <a href="/admin/education/subjects/{{ .subject.GetID }}/advance-options/modify-static-description" class="btn btn-primary">Modify static description</a>
            </div>
          </div>
          <div>
            <h2>Modify details</h2>
            <div class="well">
              You can use this option to modify the subject. <br/>
              <a href="/admin/education/subjects/{{ .subject.GetID }}/advance-options/modify" class="btn btn-primary">Modify subject</a>
            </div>
          </div>
          <div>
            <h2>Delete</h2>
            <div class="well">
              <h5>
                <span class="text-warning">
                  <span class="glyphicon glyphicon-warning-sign"></span> Please note that the data related to <strong>{{ .subject.GetName }}</strong> will be removed and it can not be recovered.
                </span>
              </h5>
              <br />
              <a data-id="{{ .subject.GetID }}" data-name="{{ .subject.GetName }}" href="#" class="btn btn-danger deleteSubjectButton">
                Delete subject from Wisply
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script src="/static/js/admin/education/subject/advance-options.js"></script>
  <script>
  $(document).ready(function(){
      var module = wisply.getModule("admin-education-subject-advance-options"),
        manager = new module.Manager();
        manager.init();
  });
  </script>
