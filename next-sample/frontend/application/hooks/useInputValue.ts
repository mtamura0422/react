import React, { useState, useCallback } from 'react';

export function useInputValue(initValue: string): [
  string,
  (e: React.ChangeEvent<HTMLInputElement>) => void, 
  React.Dispatch<React.SetStateAction<string>>
] {
  const [value, setValue] = useState(initValue);

  const updateValue = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      console.log("event.target=" + event.target.value);
      setValue(event.target.value.trimStart());
    },
    [setValue]
  );
  
  return [value, updateValue, setValue];
}


export function useInputContent(initValue: string): [
  string,
  (e: React.ChangeEvent<HTMLTextAreaElement>) => void
] {
  const [value, setValue] = useState(initValue);

  const updateValue = useCallback(
    (event: React.ChangeEvent<HTMLTextAreaElement>) => {
      setValue(event.target.value.trimStart());
    },
    [setValue]
  );
  
  return [value, updateValue];
}

