import React, { useCallback } from "react";
import CodeMirror from "@uiw/react-codemirror";
import { monokai } from "@uiw/codemirror-theme-monokai";
import { cpp } from "@codemirror/lang-cpp";

function DeafaultEditor({ setCppCode }) {
  const onChange = useCallback((value, viewUpdate) => {
    setCppCode(value);
    localStorage.setItem("cppCode", value);
  }, []);
  return (
    <div>
      <div>
        <CodeMirror
          value={
            typeof window !== "undefined" &&
              localStorage.getItem("cppCode") !== null
              ? localStorage.getItem("cppCode")
              : ""
          }
          height="90vh"
          width="auto"
          theme={monokai}
          extensions={[cpp()]}
          onChange={onChange}
        />
      </div>
    </div>
  );
}

export default DeafaultEditor;
