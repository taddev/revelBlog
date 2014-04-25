$('#cancel').click(function(){

  // redirect
  window.location.href = '/admin/';

  // prevent default behavior of button
  return false;
});


alertTimeout = setTimeout(function() {
	$('.alert-success').fadeOut("slow", "linear");
}, 3000);

var opts ={
container: 'epiceditor',
  textarea: 'PostBody',
  basePath: 'epiceditor',
  clientSideStorage: true,
  localStorageName: 'epiceditor',
  useNativeFullscreen: true,
  parser: marked,
  file: {
    name: 'epiceditor',
    defaultContent: '',
    autoSave: 100
  },
  theme: {
    base: '/public/themes/base/epiceditor.css',
    preview: '/public/themes/preview/preview-dark.css',
    editor: '/public/themes/editor/epic-dark.css'
  },
  button: {
    preview: true,
    fullscreen: true,
    bar: "auto"
  },
  focusOnLoad: false,
  shortcut: {
    modifier: 18,
    fullscreen: 70,
    preview: 80
  },
  string: {
    togglePreview: 'Toggle Preview Mode',
    toggleEdit: 'Toggle Edit Mode',
    toggleFullscreen: 'Enter Fullscreen'
  },
  autogrow: false
}
}

var editor = new EpicEditor().load(opts);

/*
$('.alert').click(function(){
	$(this).fadeOut("slow", "linear");
	clearTimeout(alertTimeout);
});
*/
