#include <iostream>
#include <string>
#include <iomanip>
using namespace std;


//swapping values
void swap(int *xp, int *yp) 
{ 
    int temp = *xp; 
    *xp = *yp; 
    *yp = temp; 
} 

//swapping values
void swapScore(double *xp, double *yp) 
{ 
    double temp = *xp; 
    *xp = *yp; 
    *yp = temp; 
} 

    //Function to  sort scores
void sortScore(double arr[], int n) 
{ 
    int i, j; 
    for (i = 0; i < n-1; i++)     
      
    // Last i elements are already in place 
    for (j = 0; j < n-i-1; j++) 
        if (arr[j] < arr[j+1]) 
            swapScore(&arr[j], &arr[j+1]); 
} 


    //Function to  sort 
void sort(int arr[], int n) 
{ 
    int i, j; 
    for (i = 0; i < n-1; i++)     
      
    // Last i elements are already in place 
    for (j = 0; j < n-i-1; j++) 
        if (arr[j] < arr[j+1]) 
            swap(&arr[j], &arr[j+1]); 
} 
  
// Function to print an array 
void printArray(int arr[], int size) { 
    for ( int i = 0; i < size; i++) 
        cout << arr[i] << " "; 
    cout << endl; 
}

// Function to calculate mean 
double getMean(double arr[], int size)  { 
    int sum  = 0;
    for ( int j = 0; j < size; j++ ) {
     sum += arr[j]; 
    //update with comments 
    }
    double mean = sum/size;
    return mean; 
}


void calculate(double arr1[],char playerIds[] ,int size)  { 

  double mean = getMean(arr1,size);
  double mode;
  double arr[size];
  for(int i = 0; i < size; i++){
      arr[i] = arr1[i];
    }

    sortScore(arr, size); 
    // we create set of arrays with sorted scores 
    //and no of players having the same score 
    double playerScores[size];
    int noPlayersWithSameScore[size]; //highest score in this array is mode
    
    int i = 0;
    int k = 0;
    int startIndex = 0;

    while(i < size) {
        if(i != size-1 && arr[i] == arr[i+1])
        {
                i++;
                continue;
        } else {
             playerScores[k] = arr[i];
             noPlayersWithSameScore[k] = i+ 1 - startIndex;
             startIndex = i+1;
             k++;
             i++;
        }

    }

    double winnerScore =  playerScores[0];
    double secondScore =  playerScores[1];
    double thirdScore =  playerScores[2];


    int noPlayersWithSameScore2[size];
  for(int i = 0; i < size; i++){
      noPlayersWithSameScore2[i] = noPlayersWithSameScore[i];
    }

    sort(noPlayersWithSameScore2, k); 
    int mostCommonPlayerScore = noPlayersWithSameScore2[0];

    if(mostCommonPlayerScore==1){ //if all repeated once, then there is no mode
        mode = -1.0; // hence we asssign -1
    }

    int modesIndex[k];
    int nrOfmodes = 0;

 for(int i = 0; i < k; i++){
      if(mostCommonPlayerScore == noPlayersWithSameScore[i]){
          modesIndex[nrOfmodes] = i;
          nrOfmodes++;
      }
    }

   double finalModes[nrOfmodes];
   for(int i = 0; i < nrOfmodes; i++){
     finalModes[i] =playerScores[modesIndex[i]];
    }

    mode = finalModes[0];

cout << "ID" << setw(10) << "Score" << setw(15) << "Results" << endl;
   string result;
   double score;
   for(int i = 0; i < size; i++){
    score = arr1[i];
     if(score == winnerScore){
         result ="Winner       ";
     }else if(score == secondScore){
         result ="Second Place ";
     }else if(score == thirdScore){
         result ="Third Place  ";
     }else if(score > mean){
         result ="Above Average";
     }else if(score < mean){
         result ="Below Average";
     }else if(score == mean){
         result ="Average      ";
     }
      cout <<  playerIds[i]<<setw(7)<< score<<setw(25)<< result<<endl; 
    //update with comments 
    }

    cout << "The mean is " << mean << ", and the mode of these values is " << mode <<"."<< endl;
  

}


int main() {
    //step1: declaring parallel arrays with logical names
    char playerIds[15] ;
    double playerScores[15] ;

    //step2: store input in arrays
    for(int i = 0; i < 15; i++){
        cin >> playerIds[i] >> playerScores[i];
    }

   //step3: print table format
    cout << "ID" << setw(10) << "Score" << endl;                     
    for ( int j = 0; j < 15; j++ ) {
      cout <<  playerIds[j]<<setw(7)<< playerScores[j]<<endl; 
    //update with comments 
    }
  
  calculate(playerScores,playerIds,15);
 
   
}