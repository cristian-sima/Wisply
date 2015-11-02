<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/repositories">Repositories</a></li>
      <li><a href="/admin/repositories/{{ .repository.ID }}">{{ .repository.Name }}</a></li>
      <li><a href="/admin/repositories/{{ .repository.ID }}/advance-options">Advance options</a></li>
      <li class="active">Modify filter</li>
    </ul>
  </div>
  <div class="panel-body">
    The filter can be used to reject or allow different types of resources, formats or categories. <br />
    <br />
    The filter should be in JSON format and it must be like this:<br />
    <ul>
      <li><strong>Operation</strong> (e.g. harvest)
        <ul>
          <li><strong>Item</strong> (e.g. records)
            <ul>
              <li>
                <strong>Action</strong> (e.g. reject) :  <strong>regex</strong> (e.g "0-9")
              </li>
            </ul>
          </li>
        </ul>
      </li>
    </ul>
    <div>
      <span class="text-warning"><span class="glyphicon glyphicon-alert"></span> Notice</span> In case your filter does not follow the format presented above, it will be ignored.
    </div>
    <br />
    <br />
    In case you are not familiar with regex, you may find this <a href="http://www.regexr.com/" target="_blank">website</a> very good.
    <br />
  </div>
  <form method="POST" id="filter-form" class="form-horizontal" >
    <input type="hidden" id="repository-filter" name="repository-filter" value=""/>
    {{ .xsrf_input }}
    <div id="jsoneditor"></div>
    <Br />
    <fieldset>
      <div class="form-group text-center">
        <div class="col-lg-10 col-lg-offset-2">
          <input type="submit" class="btn btn-primary" value="Change" /> <a class="btn btn-primary" href="/admin/repositories/{{ .repository.ID }}/advance-options">Cancel</a>
        </div>
      </div>
    </fieldset>
  </form>
</div>
<link href="/static/3rd_party/product/jsoneditor/jsoneditor.min.css" type="text/css" rel="stylesheet" property='stylesheet' />
<script src="/static/3rd_party/product/jsoneditor/jsoneditor.min.js"></script>
<script>
// Stores the JSON editor
var editor;
$(document).ready(function() {
  $("#repository-filter").focus();
  $("#filter-form").submit(formSubmited);
  loadEditor();
});
/**
* It is called when the form is sent. It takes the value of the filter and sends the form
* @param  {event} event The event which is geneated
*/
function formSubmited(event) {
  // get json
  var json = JSON.stringify(editor.get());
  $("#repository-filter").val(json);
}
/**
* It loads the filter from database into the editor
*/
function loadEditor() {
  var container = document.getElementById('jsoneditor'),
  raw = "{{ .repository.GetFilter }}";
  editor = new JSONEditor(container);
  if (raw !== "") {
    var json = JSON.parse(raw);
    editor.set(json);
    editor.expandAll();
  }
}
</script>
