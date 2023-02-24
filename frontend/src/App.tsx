import React from "react";
import axios from "axios";
import { useState } from "react";

type Winner = {
  uid: Number;
  name: String;
};

function LotteryForm({
  setWinners,
}: {
  setWinners: (winners: Winner[]) => void;
}) {
  const [url, setURL] = useState("");
  const [winnerCount, setWinnerCount] = useState(1);

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    axios
      .post(`${process.env.REACT_APP_BACKEND_URL}/api/lottery`, {
        url: url,
        winnerCount: winnerCount,
      })
      .then((resp) => {
        setWinners(resp.data.data);
      })
      .catch((reason) => console.log("Failed to get lottery:", reason));
  };

  return (
    <form
      className="space-y-8 divide-y divide-gray-200"
      onSubmit={handleSubmit}
    >
      <div className="space-y-8 divide-y divide-gray-200">
        <div className="space-y-6">
          <div>
            <h3 className="text-lg font-medium leading-6 text-gray-900">
              Bilibili 抽奖小程序
            </h3>
            <p className="mt-1 max-w-2xl text-sm text-gray-500">
              非官方抽奖小程序，对抽奖结果概不负责！
            </p>
          </div>

          <div className="space-y-6 sm:space-y-5">
            <div className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:border-t sm:border-gray-200 sm:pt-5">
              <label
                htmlFor="url"
                className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2"
              >
                视频或动态 URL
              </label>
              <div className="mt-1 col-span-2">
                <input
                  type="url"
                  name="url"
                  id="url"
                  placeholder="https://t.bilibili.com/763258243266904116"
                  value={url}
                  onChange={(event) => setURL(event.target.value)}
                  className="block w-full max-w-lg rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                />
              </div>
            </div>

            <div className="sm:grid sm:grid-cols-3 sm:items-start sm:gap-4 sm:border-t sm:border-gray-200 sm:pt-5">
              <label
                htmlFor="winners"
                className="block text-sm font-medium text-gray-700 sm:mt-px sm:pt-2"
              >
                中奖数量
              </label>
              <div className="mt-1 col-span-2">
                <input
                  type="number"
                  name="winners"
                  id="winners"
                  value={winnerCount}
                  onChange={(event) => {
                    const winnerCount = Number(event.target.value);
                    setWinnerCount(winnerCount <= 0 ? 1 : winnerCount);
                  }}
                  className="block rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
                />
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="pt-5">
        <div className="flex justify-end">
          <button
            type="submit"
            className="ml-3 inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          >
            开始抽奖
          </button>
        </div>
      </div>
    </form>
  );
}

function LotteryResult({
  winners,
  setWinners,
}: {
  winners: Winner[];
  setWinners: (winners: Winner[]) => void;
}) {
  return (
    <div className="space-y-8 divide-y divide-gray-200">
      <div className="px-6 lg:px-8">
        <div className="mt-8 flow-root">
          <div className="-my-2 -mx-6 overflow-x-auto lg:-mx-8">
            <div className="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
              <h3 className="text-lg font-medium leading-6 text-gray-900">
                中奖名单
              </h3>
              <table className="min-w-full divide-y divide-gray-300">
                <thead>
                  <tr>
                    <th
                      scope="col"
                      className="py-3.5 pl-6 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0"
                    >
                      ID
                    </th>
                    <th
                      scope="col"
                      className="py-3.5 px-3 text-left text-sm font-semibold text-gray-900"
                    >
                      昵称
                    </th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-200">
                  {winners.map((winner) => (
                    <tr key={String(winner.uid)}>
                      <td className="whitespace-nowrap py-4 pl-6 pr-3 text-sm font-medium text-gray-900 sm:pl-0">
                        {String(winner.uid)}
                      </td>
                      <td className="whitespace-nowrap py-4 px-3 text-sm text-gray-500">
                        {winner.name}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
              <button
                type="button"
                className="rounded-md border border-gray-300 bg-white py-2 px-4 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
                onClick={() => {
                  setWinners([]);
                }}
              >
                重新开始
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

function Lottery() {
  const [winners, setWinners] = useState<Winner[]>([]);

  return (
    <div className="mt-8 mx-auto max-w-xl">
      {winners.length === 0 ? (
        <LotteryForm setWinners={setWinners} />
      ) : (
        <LotteryResult winners={winners} setWinners={setWinners} />
      )}
    </div>
  );
}

export default function App() {
  return <Lottery />;
}
