#include <stdio.h>
#include <time.h>

#define EPOCH 1356076800 /* December 21, 2012 */
#define BAKTUN 13
#define SECONDS_IN_A_DAY 86400

int main() {
	printf("%d.", BAKTUN);
	int seconds_since_epoch = (int)time(NULL) - EPOCH;
	int days_since_epoch = seconds_since_epoch / SECONDS_IN_A_DAY;
	int days_in_cycle[4] = {7200, 360, 20, 1};
	for (int i = 0; i < 4; i++) { /* iterate thru each cycle */
		int cycles = days_since_epoch / days_in_cycle[i];
		days_since_epoch -= cycles * days_in_cycle[i];
		printf("%d", cycles);
		printf(i != 3 ? "." : "\n");
	}
	return 0;
}