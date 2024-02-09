from typing import List


class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [float("inf")] * (
            amount + 1
        )  # amount + 1 because we want to be able to do dp[7]
        dp[0] = 0  # Base case

        # Range from all the posible amount
        for a in range(1, amount + 1):
            # For each coin
            for c in coins:
                # If the coin isn't big enough
                if a - c >= 0:
                    # choose between the actual minimum or a new
                    dp[a] = min(dp[a], 1 + dp[a - c])

        return dp[amount] if dp[amount] != amount + 1 else -1
