<div class="col-lg-8 col-md-7 col-sm-6" >
    <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
            <ul class="breadcrumb">
                <li><a href="/admin">Admin</a></li>
                <li><a href="/admin/source">Sources</a></li>
                <li class="active">Add</li>
            </ul></div>
        <div class="panel-body">
            <p>
            <form action="/admin/source/add" method="POST" class="form-horizontal" >
                <fieldset>
                    <legend>Add a source</legend>
                    <div class="form-group">
                        <label for="source-name" class="col-lg-2 control-label">Name</label>
                        <div class="col-lg-10">
                            <input type="text" class="form-control" name="source-name" id="source-name" placeholder="Name" required>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="source-URL" class="col-lg-2 control-label">URL</label>
                        <div class="col-lg-10">
                            <input type="url" class="form-control" name="source-URL" id="source-URL" placeholder="URL" required>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="source-description" class="col-lg-2 control-label">Description</label>
                        <div class="col-lg-10">
                            <textarea class="form-control" rows="3" name="source-description" id="source-description"></textarea>
                            <span class="help-block">This field may contain notes or the description of the intitution.</span>
                        </div>
                    </div>
                    
                    <div class="form-group">
                        <div class="col-lg-10 col-lg-offset-2">
                            <a href="/admin/source"> <button type="reset" class="btn btn-default">Reset</button></a>
                            <button type="submit" class="btn btn-primary">Submit</button>
                        </div>
                    </div>
                </fieldset>
            </form>
            </p>
        </div>
    </div>
</div>

<script>
$(document).ready(function() {
    $("#source-name").focus();
});
</script>