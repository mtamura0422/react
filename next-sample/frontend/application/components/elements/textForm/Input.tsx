import React from 'react';

type TextFormProps = {
  placeholder?: string;
  value?: string | number;
  type?: string;
  min?: number;
  max?: number;
  maxLength?: number;
  required?: boolean;
  disabled?: boolean;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const Input = (props: TextFormProps) => {
  const { placeholder, min, max } = props;
  console.log('!!!!rendering InputWithte!!!!')
  return (

    
    <input
      placeholder={placeholder}
      value={props.value} // ここにvalueプロパティを追加（編集ページの初期値表示のために必要）
      type={props.type}
      min={props.type === 'number' ? min : undefined}
      max={props.type === 'number' ? max : undefined}
      // type={props.type}がnumberの場合はtext-centerを追加
      className={`${
        props.type === 'number' ? 'pl-10' : ''
      } peer h-full w-full rounded-none border-b bg-transparent 
      pb-2 pt-7 text-base caret-orange-500 
      transition-all focus:border-orange-500 focus:outline-0`}
      onChange={props.onChange}
      maxLength={props.maxLength}
      required={props.required}
      disabled={props.disabled}
    />

  );
};

export default Input;