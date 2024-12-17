import React from "react";
import ReactQuill from "react-quill";
import { useState ,useRef} from "react";
import "react-quill/dist/quill.snow.css";

const WriteBlogPage = () => {
  const [editorValue, setEditorValue] = useState("");
  const quillRef = useRef(null);

  const handleUndo =()=> {
    if(quillRef.current) {
        console.log("Undo !")
        quillRef.current.getEditor().history.undo();
    }
  };


  const handleRedo=()=>{
    if(quillRef.current) {
        console.log("Redo !")
        quillRef.current.getEditor().history.redo();
    }
  }
  const handleChange = (value) => {
    console.log(value)
    setEditorValue(value);
  };

  return (
    <div className="max-w-3xl lg:max-w-7xl mx-auto p-6 bg-white shadow-lg rounded-lg mt-10 glowing-border">
      <h2 className="text-2xl font-semibold text-gray-700 mb-4 logo-font">
        Kraft your Experience !
      </h2>
      <ReactQuill
        value={editorValue}
        onChange={handleChange}
        theme="snow"
        ref={quillRef}
        placeholder="Compose an epic..."
        modules={{
          toolbar: [
            [{ header: "1" }, { header: "2" },{header:"3"}, { font: [] }],
            [{ list: "ordered" }, { list: "bullet" }],
            [{ align: [] }],
            ["bold", "italic", "underline", "strike"],
            ["link"],
            ["image"],
            ["blockquote"],
            ["code-block"],
            ["video"],
            ["clean"],
          ],
        }}
        formats={[
          "header",
          "font",
          "list",
          "align",
          "bold",
          "strike",
          "ordered",
          "italic",
          "underline",
          "link",
          "image",
          "blockquote",
          "code-block",
          "video",
          "clean",
        ]}
        className="text-editor"
      />
      <div className="mt-6 flex flex-wrap justify-end space-x-4">
        <div className="z-50">
          <button 
          
          className="bg-blue-500 text-white px-6 py-2 rounded-lg hover:bg-blue-600 transition-colors">
            Save Draft
          </button>
        </div>

        <div className="space-x-2 z-50">
        <button
          onClick={handleUndo}
          className="bg-gray-300 text-gray-700 px-6 py-2 rounded-3xl hover:bg-gray-400  border-slate-600 border-4 transition-colors"
        >
          Undo
        </button>
        <button
          onClick={handleRedo}
          className="bg-gray-300 text-gray-700 px-6 py-2 rounded-3xl border-slate-600 border-4 hover:bg-gray-400 transition-colors"
        >
          Redo
        </button>
        </div>

        <div className="z-50">
          <button className="bg-yellow-400 hover:bg-yellow-500 text-white px-6 py-2  rounded-lg transition-colors">
            Publish
          </button>
        </div>
      </div>
    </div>
  );
};

export default WriteBlogPage;
