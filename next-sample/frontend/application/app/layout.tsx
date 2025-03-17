import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";
import Header from '@/components/layouts/header/Header';
import Footer from '@/components/layouts/footer/Footer';
import { RecoilRoot } from 'recoil';

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "離乳食レシピ",
  description: "離乳食レシピを検索できます",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >

      <div className="bg-red-50 text-[#87776c]">
        {/* ↓ここにボーダーを入れるとフッターアイコンがズレるので注意 */}
        <div className="mx-auto flex h-[100svh]  flex-col bg-white outline outline-2 outline-[#f9f2e8] lg:w-[473px] ">
          <Header />
          <div className="flex-1 overflow-y-auto overscroll-none">
              {children}
          </div>
          <Footer />
        </div>
      </div>
      </body>
    </html>
  );
}
