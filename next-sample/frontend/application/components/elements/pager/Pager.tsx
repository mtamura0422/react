"use client"

import { useRouter, useSearchParams } from 'next/navigation';

export type PageData = {
  page: number;
  q?: string;
  total_count: number;
  per_page: number;
  page_path: string;
};

export const Pagination = (props: PageData) => {

  const router = useRouter();
  var { page, q, total_count, per_page, page_path } = props;
  
  var total_page = Math.ceil(total_count / per_page)
  page = Number(page);

  const range = (start:number, end:number) =>
    [...Array(end - start + 1)].map((_, i) => start + i)

  const handlePageChange = (page: number) => {
    var pathname = `${page_path}?page=${page}`
    if (q != "") {
      pathname = `${pathname}&q=${q}`
    }
    router.push(pathname);
  };


  return (
<>
    {total_page > 1 && (
      <div className="flex justify-center w-full max-w-5xl mt-6 space-x-2">
        {page > 1 && (
          <button
          onClick={() => handlePageChange(page - 1)}
            className="bg-gray-300 text-black px-4 py-2 sm:px-6 sm:py-3 rounded text-sm sm:text-base w-20 h-10 sm:w-24 sm:h-12 hover:bg-gray-400"
          >
            前へ
          </button>
        )}
        
        <div className="flex justify-center space-x-2">
        {range(1, total_page).map((num, index) => (
          <button
                  key={index}
                  onClick={() => typeof num === 'number' && handlePageChange(num)}
                  className={`w-10 h-10 sm:w-12 sm:h-12 mx-1 rounded text-sm sm:text-base ${page === num ? 'bg-blue-500 text-white' : 'bg-gray-300 text-black hover:bg-gray-400'}`}
                  disabled={typeof num !== 'number'}
                >
                  {num}
                </button>
))}
        </div>
        {page < total_page && (
          <button
            onClick={() => handlePageChange(page + 1)}
            className="bg-gray-300 text-black px-4 py-2 sm:px-6 sm:py-3 rounded text-sm sm:text-base w-20 h-10 sm:w-24 sm:h-12 hover:bg-gray-400"
          >
            次へ
          </button>
        )}
      </div>
    )}
    </>

  );
};

export default Pagination;