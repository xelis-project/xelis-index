export type Json =
  | string
  | number
  | boolean
  | null
  | { [key: string]: Json }
  | Json[]

export interface Database {
  public: {
    Tables: {
      blocks: {
        Row: {
          block_type: string | null
          cumulative_difficulty: number | null
          difficulty: number | null
          hash: string
          height: number | null
          miner: string | null
          nonce: number | null
          reward: number | null
          size: number | null
          supply: number | null
          timestamp: number | null
          tips: string[] | null
          topoheight: number | null
          total_fees: number | null
        }
        Insert: {
          block_type?: string | null
          cumulative_difficulty?: number | null
          difficulty?: number | null
          hash: string
          height?: number | null
          miner?: string | null
          nonce?: number | null
          reward?: number | null
          size?: number | null
          supply?: number | null
          timestamp?: number | null
          tips?: string[] | null
          topoheight?: number | null
          total_fees?: number | null
        }
        Update: {
          block_type?: string | null
          cumulative_difficulty?: number | null
          difficulty?: number | null
          hash?: string
          height?: number | null
          miner?: string | null
          nonce?: number | null
          reward?: number | null
          size?: number | null
          supply?: number | null
          timestamp?: number | null
          tips?: string[] | null
          topoheight?: number | null
          total_fees?: number | null
        }
      }
      transaction_blocks: {
        Row: {
          block_hash: string | null
          tx_hash: string
        }
        Insert: {
          block_hash?: string | null
          tx_hash: string
        }
        Update: {
          block_hash?: string | null
          tx_hash?: string
        }
      }
      transaction_transfers: {
        Row: {
          amount: number | null
          asset: string | null
          extra_data: Json | null
          index: number
          to: string | null
          tx_hash: string
        }
        Insert: {
          amount?: number | null
          asset?: string | null
          extra_data?: Json | null
          index: number
          to?: string | null
          tx_hash: string
        }
        Update: {
          amount?: number | null
          asset?: string | null
          extra_data?: Json | null
          index?: number
          to?: string | null
          tx_hash?: string
        }
      }
      transactions: {
        Row: {
          fee: number | null
          hash: string
          nonce: number | null
          owner: string | null
          signature: string | null
        }
        Insert: {
          fee?: number | null
          hash: string
          nonce?: number | null
          owner?: string | null
          signature?: string | null
        }
        Update: {
          fee?: number | null
          hash?: string
          nonce?: number | null
          owner?: string | null
          signature?: string | null
        }
      }
    }
    Views: {
      [_ in never]: never
    }
    Functions: {
      get_stats: {
        Args: {
          interval: string
        }
        Returns: {
          time: string
          avg_difficulty: number
          sum_size: number
          avg_block_size: number
          count_tx: number
          sum_supply: number
          avg_block_time: number
          sum_block_fees: number
          avg_block_fees: number
          sum_block_reward: number
          avg_block_reward: number
        }[]
      }
    }
    Enums: {
      [_ in never]: never
    }
    CompositeTypes: {
      [_ in never]: never
    }
  }
}

