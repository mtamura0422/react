import React, { useCallback, useEffect, useRef, useState } from 'react';
import SearchForm from "@/features/recipe/components/Form/SearchForm"
import Result from "@/features/recipe/components/Search/Result"


export default async function Page({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string | string[] | number | undefined }>
}) {
  var { page = '1', q = '' } = await searchParams
 

  // この下からリターンの中身
    return (
      <div className="select-none">
        <div className="py-5 px-5">
        <div className="flex items-center ">
          <SearchForm
            placeholder="キーワードから探す"
            width="w-full"
            label="検索"
        
          />
          </div>
        </div>
        {(typeof q === "string") && (q != "") &&(typeof page === "string")? (
          <Result
            q={q}
            page={page}
          />
        ) : null}
      </div>
    );

};
