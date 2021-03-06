{{ define "javascript" }}

<script src="/static/3rd_party/product/paper/jquery-1.10.2.min.js"></script>
<script src="/static/3rd_party/product/paper/bootstrap.min.js"></script>
<script src="/static/3rd_party/product/paper/bootswatch.js"></script>
<script src="/static/3rd_party/product/paper/paper.js"></script>

<script src="/static/3rd_party/others/js/bootbox.min.js"></script>
<script src="/static/3rd_party/others/js/base64_decode.js"></script>
<script src="/static/3rd_party/others/js/jquery.cookie.js"></script>
<script src="/static/3rd_party/others/js/jquery.hotkeys.js"></script>

<script src="/static/3rd_party/others/js/handlebars-v4.0.2.js"></script>
<script src="/static/3rd_party/others/js/typeahead.bundle.js"></script>

<script src="/static/js/wisply/wisply.js"></script>

<script src="/static/js/wisply/search.js"></script>

{{ if .accountConnected }}
<script src="/static/js/wisply/connection.js"></script>
{{end}}

<script>
    window.cookieconsent_options = {"message":"Wisply uses cookies to ensure you get the best experience","dismiss":"Got it!","learnMore":"More info","link":"/cookies","theme":"light-bottom"};
</script>
<script src="/static/3rd_party/others/js/cookieconsent.latest.min.js"></script>

<!-- Google Analytics-->
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
    (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
    m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
  ga('create', 'UA-67059306-1', 'auto');
  ga('send', 'pageview');
</script>

<script>
$(document).ready(function(){
  var module = wisply.getModule("search");
  wisply.search = new module.Field({
    selector: '.wisply-search-field',
    URL: "/developer/api/search/anything/",
    saveSearches: true,
  });
});
</script>
{{ end }}
