import React from "react";
import CodeMirror from "@uiw/react-codemirror";
import { monokai } from "@uiw/codemirror-theme-monokai";

function Editor({ onChange, value, language, type }) {
  return (
    <div>
      <CodeMirror
        value={
          typeof window !== "undefined" && localStorage.getItem(type) !== null
            ? localStorage.getItem(type)
            : value
        }
        height="90vh"
        width="auto"
        theme={monokai}
        extensions={[language]}
        onChange={onChange}
      />
    </div>
  );
}

export default Editor;
