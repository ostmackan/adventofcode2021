use std::fs::File;
use std::path::Path;
use std::io::{self, BufRead};


struct BingoBoard {
    board: [[i32; 5 ] ; 5],
    marked: [[i32; 5] ; 5],
    bingo: bool
}

trait BingoRules {
    fn new(values: [[i32; 5]; 5]) -> Self;
    fn mark(&mut self, value: i32) -> bool;
    fn is_bingo(&mut self) -> bool;
    fn check_score(&self) -> i32;
}

impl BingoRules for BingoBoard{
    fn new(values: [[i32; 5]; 5])-> BingoBoard{
        BingoBoard {board: values, marked: [[-1; 5];5], bingo: false}
    }

    fn mark(&mut self, value: i32) -> bool{
        for x in 0..5{
            for y in 0..5{
                if self.board[x][y] == value{
                    self.marked[x][y] = value;
                    self.board[x][y] = -1;
                    println!("marked {}", value);
                    return true;
                }
            }
        }
        false
    }

    fn is_bingo(&mut self) -> bool{
        let mut bingo: bool  = false;
        //rows
        for x in 0..5{
            for y in 0..5{
                if self.marked[x][y] == -1{
                    bingo = false;
                    break;
                }else{
                    bingo = true;
                }
            }
            if bingo == true{
                break;
            }
        }
        if bingo {
            self.bingo = true;
            return true;
        }
        //columns
        for y in 0..5{
            for x in 0..5{
                if self.marked[x][y] == -1{
                    bingo = false;
                    break;
                }else{
                    bingo = true;
                }
            }
            if bingo == true{
                break;
            }
        }
        if bingo {
            self.bingo = true;
            return true;
        } else{
            return false;
        }
    }

    fn check_score(&self) -> i32{
        let mut result: i32 = 0;
        for x in 0..5{
            for y in 0..5{
                if self.board[x][y] != -1{
                    result = result + self.board[x][y];
                }
            }
        }

        return result;
    }
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}

fn run(filename: &str){
    let mut called: Vec<i32> = Vec::new();
    let mut boards: Vec<BingoBoard> = Vec::new();
    let mut board: [[i32; 5 ] ; 5] = [[-1; 5 ] ; 5];
    let mut x: usize = 0;
    let mut y: usize = 0;
    if let Ok(lines) = read_lines(filename){
        for line in lines{
            if let Ok(ip) = line {
                //println!("{}", ip);
                if ip.contains(",") == true{
                    //first row
                    let v: Vec<&str> = ip.split(',').collect();
                    
                    for value in v.iter(){
                        let c: i32 = value.parse().unwrap();
                        called.push(c);
                    }

                    x = 0;
                    y = 0;
                }
                else if ip.len() > 1{
                    let v: Vec<&str> = ip.split(' ').collect();
                    for value in v.iter(){
                        if value.trim().is_empty() == false{
                            let c: i32 = value.parse().unwrap();
                            board[x][y] = c;
                            y = y+1;
                        }
                    }
                    x = x + 1;
                    if x > 4{
                        boards.push(BingoRules::new(board));
                    }
                    y = 0;
                }
                else {
                    //end a board
                    
                    board = [[-1; 5 ] ; 5];
                    x = 0;

                }
            }
        }
    }

    for active_board in boards.iter(){
        //println!("{} {} {} {} {}", active_board.board[0][0], active_board.board[0][1], active_board.board[0][2], active_board.board[0][3],active_board.board[0][4] );
        //println!("{} {} {} {} {}", active_board.marked[0][0], active_board.marked[0][1], active_board.marked[0][2], active_board.marked[0][3],active_board.marked[0][4] );
    }

    println!(" ");

    let mut boards_without_bingo: usize = boards.len()+1;
    let mut bingo: bool = false;
    for call in called.iter(){
        for active_board in boards.iter_mut(){
            if active_board.bingo != true{
            if active_board.mark(*call) == true {
                if active_board.is_bingo(){
                    println!("BINGO {} {} = {} * {}",boards_without_bingo, active_board.check_score() * call, active_board.check_score(), call);
                    println!("{} {} {} {} {}", active_board.board[0][0], active_board.board[0][1], active_board.board[0][2], active_board.board[0][3],active_board.board[0][4] );
                    println!("{} {} {} {} {}", active_board.board[1][0], active_board.board[1][1], active_board.board[1][2], active_board.board[1][3],active_board.board[1][4] );
                    println!("{} {} {} {} {}", active_board.board[2][0], active_board.board[2][1], active_board.board[2][2], active_board.board[2][3],active_board.board[2][4] );
                    println!("{} {} {} {} {}", active_board.board[3][0], active_board.board[3][1], active_board.board[3][2], active_board.board[3][3],active_board.board[3][4] );
                    println!("{} {} {} {} {}", active_board.board[4][0], active_board.board[4][1], active_board.board[4][2], active_board.board[4][3],active_board.board[4][4] );
                    println!(" ");
                    boards_without_bingo = boards_without_bingo-1;
                    if boards_without_bingo == 0{
                        bingo = true
                    }
                    
                }
                
            }else{
                println!("{}", call);
            }
        }
            if bingo{
                break;
            }
        }
        if bingo {
            break;
        }
    }
}

fn main() {
    run("input.txt");
}
