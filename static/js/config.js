require.config({
	baseUrl: '/static/lib',
	paths: {
		jquery          : "jquery",
		bootstrap       : "bootstrap/js/bootstrap.min",
		ckeditor        : "ckeditor/ckeditor",
        // marked          : "marked.min",
        // prettify        : "prettify.min",
        // raphael         : "raphael.min",
        // underscore      : "underscore.min",
        // flowchart       : "flowchart.min", //flowchart
        // jqueryflowchart : "jquery.flowchart.min", 
        // sequenceDiagram : "sequence-diagram.min",
        // katex           : "//cdnjs.cloudflare.com/ajax/libs/KaTeX/0.1.1/katex.min",
        // editormd        : "editor.md/editormd.amd" // Using Editor.md amd version for Require.js
	},
	shim:{
		'bootstrap': {
			exports: "$",
			deps: ['jquery']
		},
		'ckeditor': {
			exports: "CKEDITOR",
			deps: ['jquery']
		}//,
		// 'editormd':{
		// 	exports:"editormd",
		// 	deps   :['jquery','sequenceDiagram','marked','prettify','flowchart','jqueryflowchart','katex']
		// },
		// 'jqueryflowchart':{
		// 	exports:"_",
		// 	deps:['jquery']
		// },
		// 'sequenceDiagram':{
		// 	exports:"Diagram",
		// 	deps:['underscore','raphael']
		// },

	}
});