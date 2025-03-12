import React from 'react';

type TextAreaProps = {
  placeholder?: string;
  value?: string | number;
  maxLength?: number;
  onChange?: (e: React.ChangeEvent<HTMLTextAreaElement>) => void
};

const TextArea = (props: TextAreaProps) => {
  const { placeholder } = props;

  return (
    <textarea
    placeholder={placeholder}
    value={props.value} // ここにvalueプロパティを追加（編集ページの初期値表示のために必要）
    className="border-blue-gray-200 text-blue-gray-700 
    placeholder-shown:border-blue-gray-200 
    disabled:bg-blue-gray-50 peer h-full 
    min-h-[200px] w-full resize-none border-b 
    bg-transparent pb-1.5 pt-6 text-base caret-orange-500 
    transition-all focus:border-orange-500 
    focus:outline-0 disabled:resize-none"
    onChange={props.onChange}
    maxLength={props.maxLength}
  ></textarea>

  );
};

export default TextArea;