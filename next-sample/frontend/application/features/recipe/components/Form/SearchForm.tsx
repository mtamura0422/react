"use client"; 

import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { usePathname, useRouter } from 'next/navigation'
import { useRecoilState } from 'recoil';
import {  useAtom } from "jotai";

import {
  searchWordState,
} from '@/state/search';



import React, { useCallback, useEffect, useRef, useState } from 'react';

import { useInputValue } from '@/hooks/useInputValue';


type TextFormProps = {
  label?: React.ReactNode;
  placeholder?: string;
  width: string;
  value?: string | number;
  type?: string;
  min?: number;
  max?: number;
  maxLength?: number;
  required?: boolean;
  disabled?: boolean;
};




/*
export const searchWordState = atom<string[]>({
  key: 'searchWordState',
  default: [''], // 初期状態は空の配列
});
*/
const SearchForm = (props: TextFormProps) => {
;
  const [searchWord, updateSearchWord] = useAtom(searchWordState);

  const pathname = usePathname();
  const { replace } = useRouter();
  const timerId = useRef<NodeJS.Timeout | null>(null); // useRefを使ってタイマーIDを管理する。stateにすると再レンダリングされてしまうため、useRefを使う


  useEffect(() => {

    if (searchWord != "") {
      const params = new URLSearchParams();
      params.set('page', '1');
      params.set('q', searchWord);
      
      replace(`${pathname}?${params.toString()}`);
    }

  /*
    if (searchWord.trim() !== '') {
      timerId.current = setTimeout(() => {
        sendSearchRequest(); // ✅ 非同期関数を呼び出すだけにする
      }, 500);
    }
      */
  }, [searchWord]);

  /*
  useEffect(() => {
    // -------------検索ワードが空でない場合-----------------
    if (searchWord !== '') {
      console.log("searchWord="+searchWord);
      // 既にタイマーがセットされている場合は、タイマーをクリアする
      if (timerId.current !== null) {
        clearTimeout(timerId.current); // ※タイマーをクリアする理由は、連続で入力された場合に、前回のタイマーをクリアするため
      }
      // timerId.currentに、setTimeoutの返り値を代入することで、タイマーIDを管理できる
      // 0.5秒後に、sendSearchRequest()を実行する
      timerId.current = setTimeout(() => {
        sendSearchRequest(); // 検索リクエストを送信する関数
      }, 500);
    }
    // ---------------検索ワードが空の場合-------------------
    else {
      setResultRecipes([]); // 検索結果を空にする
      // 既にタイマーがセットされている場合は、タイマーをクリアする
      if (timerId.current !== null) {
        clearTimeout(timerId.current); // ※タイマーをクリアする理由は、連続で入力された場合に、前回のタイマーをクリアするため
      }
    }
  }, [ searchWord, sendSearchRequest]);
*/
  //const [searchWords, setSearchWords] = useRecoilState(searchWordState);

  const { placeholder, width, min, max } = props;
console.log("searchWord = " + searchWord)
  return (
    <div className={`flex ${width} relative`}>
      {/* relativeを追加 */}
      <input
        placeholder={placeholder}
        type={props.type}
        min={props.type === 'number' ? min : undefined}
        max={props.type === 'number' ? max : undefined}
        value={searchWord}
        className="peer h-full w-full rounded-lg bg-transparent py-1.5 pl-8 caret-pink-500 outline outline-1 outline-gray-300 transition-all duration-100 focus:outline-2 focus:outline-pink-500"
        onChange={(e) => {
          const trimmedValue = e.target.value.trimStart(); // 入力値の先頭の空白を除去
          updateSearchWord(trimmedValue); // 検索ワードをstateにセット

        }}
        maxLength={props.maxLength}
        required={props.required}
        disabled={props.disabled}
      />
      <FontAwesomeIcon
        icon={faMagnifyingGlass}
        className="absolute left-2 top-1/2 -translate-y-1/2 text-lg text-gray-400 transition-all duration-100 peer-focus:text-base peer-focus:text-pink-500"
      />

    </div>

   

  );
};
export default SearchForm;