/*
 * Project: curve
 * Created Date: 2019-11-25
 * Author: lixiaocui
 * Copyright (c) 2018 netease
 */

#ifndef SRC_TOOLS_NBD_INTERRUPTIBLE_SLEEPER_H_
#define SRC_TOOLS_NBD_INTERRUPTIBLE_SLEEPER_H_

#include <chrono>  // NOLINT
#include <condition_variable>  // NOLINT
#include <memory>
#include <mutex>  // NOLINT

namespace curve {
namespace nbd {
/**
 * InterruptibleSleeper 实现可 interruptible 的 sleep 功能.
 * 正常情况下 wait_for 超时, 接收到退出信号之后, 程序会立即被唤醒,
 * 退出 while 循环, 并执行 cleanup 代码.
 */
class InterruptibleSleeper {
 public:
    /**
     * @brief wait_for 等待指定时间，如果接受到退出信号立刻返回
     * @param[in] time 指定wait时长
     * @return false-收到退出信号 true-超时后退出
     */
    template <typename R, typename P>
    bool wait_for(std::chrono::duration<R, P> const& time) {
        std::unique_lock<std::mutex> lock(m);
        return !cv.wait_for(lock, time, [&] { return terminate; });
    }

    /**
     * @brief interrupt 给当前wait发送退出信号
     */
    void interrupt() {
        std::lock_guard<std::mutex> lock(m);
        terminate = true;
        cv.notify_all();
    }

 private:
    std::condition_variable cv;
    std::mutex m;
    bool terminate = false;
};

}  // namespace nbd
}  // namespace curve

#endif  // SRC_TOOLS_NBD_INTERRUPTIBLE_SLEEPER_H_