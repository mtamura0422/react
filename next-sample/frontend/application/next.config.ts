import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactStrictMode: false,
  images: {
    domains: ['konbini-recipe.s3.ap-northeast-1.amazonaws.com']
  }
};

export default nextConfig;
